package router

import (
	"net/http"
	"regexp"

	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products"
)

type route struct {
	method  string
	pattern *regexp.Regexp
	handler http.HandlerFunc
	params  []string
}

type Router struct {
	routes []route
}

func New(
	productsController products.Controller,
) *Router {
	router := Router{}

	router.GET("/api/v1/products", productsController.GetProducts)
	router.POST("/api/v1/products", productsController.PostProducts)
	router.GET("/api/v1/products/:productid", productsController.GetProduct)
	router.PUT("/api/v1/products/:productid", productsController.PutProduct)
	router.DELETE("/api/v1/products/:productid", productsController.DeleteProduct)

	return &router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if r.Method != route.method {
			continue
		}

		matches := route.pattern.FindStringSubmatch(r.URL.Path)

		if len(matches) > 0 {

			route.handler(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func (router *Router) addRoute(method string, pattern string, handler http.HandlerFunc) {
	paramMatcher := regexp.MustCompile(":([a-zA-Z]+)")
	paramMatches := paramMatcher.FindAllStringSubmatch(pattern, -1)

	params := make([]string, len(paramMatches))

	if len(paramMatches) > 0 {
		pattern = paramMatcher.ReplaceAllLiteralString(pattern, "([^/]+)")

		for i, match := range paramMatches {
			params[i] = match[1]
		}
	}

	router.routes = append(router.routes, route{
		method:  method,
		pattern: regexp.MustCompile("^" + pattern + "$"),
		handler: handler,
		params:  params,
	})
}

func (router *Router) GET(pattern string, handler http.HandlerFunc) {
	router.addRoute(http.MethodGet, pattern, handler)
}

func (router *Router) POST(pattern string, handler http.HandlerFunc) {
	router.addRoute(http.MethodPost, pattern, handler)
}

func (router *Router) PUT(pattern string, handler http.HandlerFunc) {
	router.addRoute(http.MethodPut, pattern, handler)
}

func (router *Router) DELETE(pattern string, handler http.HandlerFunc) {
	router.addRoute(http.MethodDelete, pattern, handler)
}
