package main

import (
	"fmt"
	"github.com/matevskial/chirpyx/handlerutils"
	"net/http"
	"sync/atomic"
)

type apiMetrics struct {
	hits atomic.Uint64
}

func (a *apiMetrics) meteredHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		a.hits.Add(1)
		next.ServeHTTP(w, req)
	})
}

func (a *apiMetrics) metricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handlerutils.RespondWithText(w, http.StatusOK, fmt.Sprintf("Hits: %v", a.hits.Load()))
	})
}

func (a *apiMetrics) metricsAdminHandler() http.Handler {
	template := `<html>
<body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
</body>
</html>
`
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handlerutils.RespondWithHtml(w, http.StatusOK, fmt.Sprintf(template, a.hits.Load()))
	})
}

func (a *apiMetrics) resetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		a.hits.Store(0)
		w.WriteHeader(http.StatusOK)
	})
}
