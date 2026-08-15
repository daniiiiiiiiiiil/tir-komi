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
	api "github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/file_handlers"
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

	postgres_media "github.com/daniiiiiiiiiiil/tir-komi/internal/features/media/repository/postgres"
	media_service "github.com/daniiiiiiiiiiil/tir-komi/internal/features/media/service"
	media_http "github.com/daniiiiiiiiiiil/tir-komi/internal/features/media/transport/http"

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

	log.Debug("Initializing feature", zap.String("feature", "auth"))
	authRepo := authRepository.NewAuthRepository(pool)
	hasher := &authService.BcryptHasher{}
	authSvc := authService.NewAuthService(authRepo, hasher)
	authController := authApi.NewAuthController(authSvc, jwtProvider)

	log.Debug("Initializing feature", zap.String("feature", "advertisement"))
	adRepo := postgres_advertisement.NewAdvertisementRepository(pool)
	adSvc := advertisementService.NewAdvertisementService(adRepo)
	adHandler := advertisementHttp.NewAdvertisementHandler(adSvc)

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
	materialHandler := methodologicalHttp.NewMaterialHandler(materialSvc)

	log.Debug("Initializing feature", zap.String("feature", "media"))
	mediaRepo := postgres_media.NewMediaRepository(pool)
	mediaSvc := media_service.NewMediaService(mediaRepo)
	mediaHandler := media_http.NewMediaHandler(mediaSvc)

	log.Debug("Initializing file handler")
	fileHandler := api.NewFileHandler(adSvc, materialSvc)

	log.Debug("Initializing feature", zap.String("feature", "web"))
	webRepo := webRepository.NewWebRepository()
	webSvc := webService.NewWebService(webRepo)
	webController := webHttp.NewWebController(webSvc)

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
	apiRouter.RegisterRouters(mediaHandler.Routers()...)

	apiRouter.RegisterRouters(server.Route{
		Method:  "GET",
		Path:    "/advertisements/{id}/image",
		Handler: fileHandler.GetAdvertisementImage,
	})
	apiRouter.RegisterRouters(server.Route{
		Method:  "GET",
		Path:    "/advertisements/{id}/pdf",
		Handler: fileHandler.GetAdvertisementPDF,
	})
	apiRouter.RegisterRouters(server.Route{
		Method:  "GET",
		Path:    "/materials/{id}/pdf",
		Handler: fileHandler.GetMaterialPDF,
	})

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
