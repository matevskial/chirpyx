package main

import (
	"log"
	"net/http"
)

func main() {
	httpServeMux := http.NewServeMux()

	/* note that if serveMux does not have any handler, the is a handler that returns 404 status code
	I added a custom handler just for demonstration purposes
	*/
	httpServeMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.NotFound(w, req)
	})

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
