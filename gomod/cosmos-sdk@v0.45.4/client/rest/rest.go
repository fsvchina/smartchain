package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)




const DeprecationURL = "https:



func addHTTPDeprecationHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Deprecation", "true")
		w.Header().Set("Link", "<"+DeprecationURL+">; rel=\"deprecation\"")
		w.Header().Set("Warning", "199 - \"this endpoint is deprecated and may not work as before, see deprecation link for more info\"")
		h.ServeHTTP(w, r)
	})
}





func WithHTTPDeprecationHeaders(r *mux.Router) *mux.Router {
	subRouter := r.NewRoute().Subrouter()
	subRouter.Use(addHTTPDeprecationHeaders)
	return subRouter
}
