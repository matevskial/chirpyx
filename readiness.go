package main

import "net/http"

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
