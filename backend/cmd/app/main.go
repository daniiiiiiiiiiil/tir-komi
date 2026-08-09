package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/config"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres/core_pgx"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/storage"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/jwt"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/middleware"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
	postgres_advertisement "github.com/daniiiiiiiiiiil/tir-komi/internal/features/advertisement/repository/postgres"
	postgres_methodological_material "github.com/daniiiiiiiiiiil/tir-komi/internal/features/methodological_material/repository/postgres"
	postgres_review "github.com/daniiiiiiiiiiil/tir-komi/internal/features/reviews/repository/postgres"
	review_service "github.com/daniiiiiiiiiiil/tir-komi/internal/features/reviews/service"
	reviews_http "github.com/daniiiiiiiiiiil/tir-komi/internal/features/reviews/transport/http"
	postgres_vacant_positions "github.com/daniiiiiiiiiiil/tir-komi/internal/features/vacant_positions/repository/postgres"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/vacant_positions/service"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/vacant_positions/transport/http"

	advertisementService "github.com/daniiiiiiiiiiil/tir-komi/internal/features/advertisement/service"
	advertisementHttp "github.com/daniiiiiiiiiiil/tir-komi/internal/features/advertisement/transport/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/auth/authRepository"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/auth/authService"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/auth/authTransport/authApi"

	methodologicalService "github.com/daniiiiiiiiiiil/tir-komi/internal/features/methodological_material/service"
	methodologicalHttp "github.com/daniiiiiiiiiiil/tir-komi/internal/features/methodological_material/transport/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/web/webRepository"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/web/webService"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/web/webTransport/webHttp"

	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log, err := logger.NewLog(logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to initialize logger, err=", err)
		os.Exit(1)
	}
	defer log.Close()

	log.Debug("Initializing postgres connection pool")
	pool, err := core_pgx.NewPool(ctx, core_pgx.NewConfigMust())
	if err != nil {
		log.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	jwtConfig := jwt.NewMustJWtConfig()
	jwtProvider := jwt.NewJwtProvider(jwtConfig)

	fileStorage := storage.NewFileStorage(storage.FileStorageConfig{
		UploadDir: "./uploads",
		MaxSize:   100 << 20,
	})
	log.Debug("File storage initialized", zap.String("upload_dir", "./uploads"))

	log.Debug("Initializing feature", zap.String("feature", "auth"))
	authRepo := authRepository.NewAuthRepository(pool)
	hasher := &authService.BcryptHasher{}
	authSvc := authService.NewAuthService(authRepo, hasher)
	authController := authApi.NewAuthController(authSvc, jwtProvider)

	log.Debug("Initializing feature", zap.String("feature", "advertisement"))
	adRepo := postgres_advertisement.NewAdvertisementRepository(pool)
	adSvc := advertisementService.NewAdvertisementService(adRepo)
	adHandler := advertisementHttp.NewAdvertisementHandler(adSvc, fileStorage)

	log.Debug("Initializing feature", zap.String("feature", "review"))
	reviewRepo := postgres_review.NewReviewRepository(pool)
	reviewSvc := review_service.NewReviewService(reviewRepo)
	reviewHandler := reviews_http.NewReviewHandler(reviewSvc)

	log.Debug("Initializing feature", zap.String("feature", "vacant_position"))
	vacantRepo := postgres_vacant_positions.NewVacantPositionRepository(pool)
	vacantSvc := service.NewVacantPositionService(vacantRepo)
	vacantHandler := http.NewVacantHandler(vacantSvc)

	log.Debug("Initializing feature", zap.String("feature", "methodological_material"))
	materialRepo := postgres_methodological_material.NewMethodologicalMaterialRepository(pool)
	materialSvc := methodologicalService.NewMethodologicalMaterialService(materialRepo)
	materialHandler := methodologicalHttp.NewMaterialHandler(materialSvc, fileStorage)

	log.Debug("Initializing feature", zap.String("feature", "web"))
	webRepo := webRepository.NewWebRepository()
	webSvc := webService.NewWebService(webRepo)
	webController := webHttp.NewWebController(webSvc)

	RegisterAssets()

	log.Debug("Initializing http server")
	httpServer := server.NewHTTPServer(
		server.NewConfigMust(),
		log,
		middleware.CORS(),
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.Trace(),
		middleware.Panic(),
	)

	apiRouter := server.NewAPIVersionRouter(server.ApiVersion1)

	apiRouter.RegisterRouters(authController.Routers()...)
	apiRouter.RegisterRouters(adHandler.Routers()...)
	apiRouter.RegisterRouters(reviewHandler.Routers()...)
	apiRouter.RegisterRouters(vacantHandler.Routers()...)
	apiRouter.RegisterRouters(materialHandler.Routers()...)

	httpServer.RegisterRoutes(webController.Routes()...)
	httpServer.RegisterAPIRouters(apiRouter)
	httpServer.RegisterSwagger()

	log.Info("Starting application...")
	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server error", zap.Error(err))
		os.Exit(1)
	}

	log.Info("Application stopped")
}
