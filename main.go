package main

import (
	"log"
	"net/http"
)

func main() {
	staticContentDir := http.Dir(".")
	httpFileServerPrefix := "/app/"
	httpFileServerMetrics := apiMetrics{}
	meteredHttpFileServer := httpFileServerMetrics.meteredHandler(http.FileServer(staticContentDir))

	httpServeMux := http.NewServeMux()

	/*
		[METHOD ][HOST]/[PATH] is the correct format of the path stribg
	*/
	httpServeMux.Handle(httpFileServerPrefix+"*", http.StripPrefix(httpFileServerPrefix, meteredHttpFileServer))
	httpServeMux.Handle("GET /api/healthz", healthHandler{})
	httpServeMux.Handle("GET /api/metrics", httpFileServerMetrics.metricsHandler())
	httpServeMux.Handle("GET /api/reset", httpFileServerMetrics.resetHandler())
	httpServeMux.Handle("GET /admin/metrics", httpFileServerMetrics.metricsAdminHandler())

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
