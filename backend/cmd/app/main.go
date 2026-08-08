package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/config"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/repository/pool/postgres/core_pgx"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/transport/http/jwt"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/transport/http/middleware"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/transport/http/server"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/features/auth/authRepository"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/features/auth/authService"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/features/auth/authTransport/authApi"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/features/web/webRepository"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/features/web/webService"
	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/features/web/webTransport/webHttp"
	"go.uber.org/zap"
)

func main() {
	config := config.NewConfigMust()
	time.Local = config.TimeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log, err := logger.NewLog(logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to initialize logger,ex=", err)
		os.Exit(1)
	}
	defer log.Close()

	log.Debug("Initialing postgres connection pool")
	pool, err := core_pgx.NewPool(ctx, core_pgx.NewConfigMust())
	if err != nil {
		log.Fatal("failed to init postgres connection poll", zap.Error(err))
	}
	defer pool.Close()

	jwtConfig := jwt.NewMustJWtConfig()
	jwtProvider := jwt.NewJwtProvider(jwtConfig)

	//txManager := core_pgx.NewPgxTxManager(pool)

	log.Debug("Initialing feature", zap.String("feature", "auth"))
	authPostgresRepository := authRepository.NewAuthRepository(pool)
	hasher := &authService.BcryptHasher{}
	authService := authService.NewAuthService(authPostgresRepository, hasher)
	authTransportHTTP := authApi.NewAuthController(authService, jwtProvider)

	log.Debug("Initialing feature", zap.String("feature", "html"))
	htmlRepository := webRepository.NewWebRepository()
	htmlService := webService.NewWebService(htmlRepository)
	htmlController := webHttp.NewWebController(htmlService)

	log.Debug("Initializing http server")
	httpServer := server.NewHTTPServer(
		server.NewConfigMust(), log,
		middleware.CORS(),
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.Trace(),
		middleware.Panic(),
	)

	//authMW := middleware.Auth(jwtProvider)
	//adminMW := middleware.RequireRole(domain.RoleAdmin)

	apiVersionRouter := server.NewAPIVersionRouter(server.ApiVersion1)

	apiVersionRouter.RegisterRouters(authTransportHTTP.Routers()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)
	httpServer.RegisterSwagger()
	httpServer.RegisterRoutes(htmlController.Routes()...)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run, err=", zap.Error(err))
	}
}
