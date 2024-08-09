package main

import (
	"log"
	"net/http"
)

func main() {
	staticContentDir := http.Dir(".")
	httpFileServer := http.FileServer(staticContentDir)

	httpServeMux := http.NewServeMux()

	httpServeMux.Handle("/", httpFileServer)

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
