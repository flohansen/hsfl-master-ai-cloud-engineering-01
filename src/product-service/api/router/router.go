package router

import (
	"net/http"

	"github.com/flohansen/hsfl-master-ai-cloud-engineering/lib/router"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products"
)

type Router struct {
	router http.Handler
}

func New(
	productsController products.Controller,
) *Router {
	router := router.New()

	router.GET("/api/v1/products", productsController.GetProducts)
	router.POST("/api/v1/products", productsController.PostProducts)
	router.GET("/api/v1/products/:productid", productsController.GetProduct)
	router.PUT("/api/v1/products/:productid", productsController.PutProduct)
	router.DELETE("/api/v1/products/:productid", productsController.DeleteProduct)

	return &Router{router}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
