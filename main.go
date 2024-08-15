package main

import (
	"github.com/matevskial/chirpyx/auth"
	"github.com/matevskial/chirpyx/configuration"
	"github.com/matevskial/chirpyx/database"
	authHandler "github.com/matevskial/chirpyx/handlers/auth"
	chirpHandler "github.com/matevskial/chirpyx/handlers/chirp"
	polkaHandler "github.com/matevskial/chirpyx/handlers/polka"
	userHandler "github.com/matevskial/chirpyx/handlers/user"
	authMiddleware "github.com/matevskial/chirpyx/middlewares/auth"
	polkaAuthMiddleware "github.com/matevskial/chirpyx/middlewares/polkaauth"
	"github.com/matevskial/chirpyx/polkaauth"
	chirpRepository "github.com/matevskial/chirpyx/repository/chirp"
	userRepository "github.com/matevskial/chirpyx/repository/user"
	"log"
	"net/http"
)

func main() {
	config, configErr := configuration.Parse()
	if configErr != nil {
		log.Fatalf("Error initializing configuration: %v", configErr)
	}

	db, dbErr := database.NewDB("database.json", config.IsDevMode)
	if dbErr != nil {
		log.Fatalf("Error initializing database: %v", dbErr)
	}

	authenticationService := auth.NewAuthenticationJwtService(config)
	refreshTokenService := auth.NewRefreshTokenService(db)
	polkaAuthenticationService := polkaauth.NewPolkaAuthenticationService(config)

	staticContentDir := http.Dir(".")
	httpFileServerPrefix := "/app/"
	httpFileServerMetrics := apiMetrics{}
	meteredHttpFileServer := httpFileServerMetrics.meteredHandler(http.FileServer(staticContentDir))

	authenticationMiddleware := authMiddleware.NewAuthenticationMiddleware(authenticationService)
	polkaAuthenticationMiddleware := polkaAuthMiddleware.NewPolkaAuthenticationMiddleware(polkaAuthenticationService)

	chirpRepo := chirpRepository.NewChirpJsonFileRepository(db)
	chirpHndlr := chirpHandler.NewChirpHandler(chirpRepo, authenticationMiddleware)

	userRepo := userRepository.NewUserJsonFileRepository(db)
	userHndlr := userHandler.NewUserHandler(userRepo, authenticationMiddleware)

	authenticationHndlr := authHandler.NewAuthenticationHandler("/api/login", userRepo, authenticationService, refreshTokenService)

	polkaHandlr := polkaHandler.NewPolkaHandler(userRepo)

	httpServeMux := http.NewServeMux()

	/*
		[METHOD ][HOST]/[PATH] is the correct format of the path string
	*/
	httpServeMux.Handle(httpFileServerPrefix+"*", http.StripPrefix(httpFileServerPrefix, meteredHttpFileServer))
	httpServeMux.Handle("GET /api/healthz", healthHandler{})
	httpServeMux.Handle("GET /api/metrics", httpFileServerMetrics.metricsHandler())
	httpServeMux.Handle("GET /api/reset", httpFileServerMetrics.resetHandler())
	httpServeMux.Handle("GET /admin/metrics", httpFileServerMetrics.metricsAdminHandler())
	httpServeMux.Handle("/api/", http.StripPrefix("/api", chirpHndlr.Handler()))
	httpServeMux.Handle("/api/users", userHndlr.Handler("/api/users"))
	httpServeMux.Handle("POST /api/login", authenticationHndlr.LoginHandler())
	httpServeMux.Handle("POST /api/refresh", authenticationHndlr.RefreshTokenHandler())
	httpServeMux.Handle("POST /api/revoke", authenticationHndlr.RevokeRefreshTokenHandler())
	httpServeMux.Handle("/api/polka/", polkaAuthenticationMiddleware.AuthenticatedHandler(polkaHandlr.Handler("/api/polka")))

	httpServer := http.Server{
		Handler: httpServeMux,
		Addr:    ":8080",
	}

	err := httpServer.ListenAndServe()
	/* note that call to log.Fatal blocks in this case, meaning  the server is running
	As of now, I am not sure why it blocks, maybe there is a channel that waits to receive an actual error value?
	*/
	log.Fatal(err)
}
