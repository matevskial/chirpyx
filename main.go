package main

import (
	"flag"
	"github.com/matevskial/chirpyx/database"
	"github.com/matevskial/chirpyx/handlers/authentication"
	chirpHandler "github.com/matevskial/chirpyx/handlers/chirp"
	userHandler "github.com/matevskial/chirpyx/handlers/user"
	chirpRepository "github.com/matevskial/chirpyx/repository/chirp"
	userRepository "github.com/matevskial/chirpyx/repository/user"
	"log"
	"net/http"
)

func main() {
	isDevMode := flag.Bool("dev", false, "Enable dev mode")
	flag.Parse()

	db, dbErr := database.NewDB("database.json", *isDevMode)
	if dbErr != nil {
		log.Fatalf("Error initializing database: %v", dbErr)
	}

	staticContentDir := http.Dir(".")
	httpFileServerPrefix := "/app/"
	httpFileServerMetrics := apiMetrics{}
	meteredHttpFileServer := httpFileServerMetrics.meteredHandler(http.FileServer(staticContentDir))

	chirpRepo := chirpRepository.NewChirpJsonFileRepository(db)
	chirpHndlr := chirpHandler.NewChirpHandler(chirpRepo)

	userRepo := userRepository.NewUserJsonFileRepository(db)
	userHndlr := userHandler.NewUserHandler("/api/users", userRepo)

	authenticationHndlr := authentication.NewAuthenticationHandler("/api/login", userRepo)

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
	httpServeMux.Handle(userHndlr.Path, userHndlr.Handler())
	httpServeMux.Handle(authenticationHndlr.Path, authenticationHndlr.Handler())

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
