package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type healthHandler struct{}

func (healthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	responseString := "OK"
	headers := w.Header()
	headers.Add("Content-Type", "text/plain; charset=utf-8")
	/* note that the line below is not really needed if w.Write is called, since 200 would be assumed
	This is used for setting custom status codes
	*/
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(responseString))
	if err != nil {
		http.Error(w, "Internal server error", 500)
	}
}

type apiMetrics struct {
	hits atomic.Uint64
}

func (a *apiMetrics) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		a.hits.Add(1)
		next.ServeHTTP(w, req)
	})
}

func (a *apiMetrics) metricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hits: %v", a.hits.Load())))
	})
}

func (a *apiMetrics) resetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		a.hits.Store(0)
		w.WriteHeader(http.StatusOK)
	})
}

func main() {
	staticContentDir := http.Dir(".")
	httpFileServerPrefix := "/app/"
	httpFileServerMetrics := apiMetrics{}
	meteredHttpFileServer := httpFileServerMetrics.middleware(http.FileServer(staticContentDir))

	httpServeMux := http.NewServeMux()

	/*
		[METHOD ][HOST]/[PATH] is the correct format of the path stribg
	*/
	httpServeMux.Handle(httpFileServerPrefix+"*", http.StripPrefix(httpFileServerPrefix, meteredHttpFileServer))
	httpServeMux.Handle("GET /healthz", healthHandler{})
	httpServeMux.Handle("GET /metrics", httpFileServerMetrics.metricsHandler())
	httpServeMux.Handle("GET /reset", httpFileServerMetrics.resetHandler())

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
