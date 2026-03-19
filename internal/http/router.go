package http

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
)
/*
create new router
add a GET /health route
when /health is called, return JSON saying status is ok
return the router
*/
func NewRouter() http.Handler {
	//new router
	r := chi.NewRouter()

	//register router
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//go map -> JSON
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})
	return r
}