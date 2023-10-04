package router

import "net/http"

type Router struct {
	mux *http.ServeMux
}

func New(
	productsHandler http.Handler,
) *Router {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/products", productsHandler)

	return &Router{mux}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}
