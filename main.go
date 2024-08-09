package main

import (
	"log"
	"net/http"
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

func main() {
	staticContentDir := http.Dir(".")
	httpFileServerPrefix := "/app/"
	httpFileServer := http.FileServer(staticContentDir)

	httpServeMux := http.NewServeMux()

	httpServeMux.Handle(httpFileServerPrefix+"*", http.StripPrefix(httpFileServerPrefix, httpFileServer))
	httpServeMux.Handle("/healthz", healthHandler{})

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
