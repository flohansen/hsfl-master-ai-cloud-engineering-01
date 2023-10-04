package products

import "net/http"

type Controller interface {
	GetProducts(http.ResponseWriter, *http.Request)
	PostProducts(http.ResponseWriter, *http.Request)
	GetProduct(http.ResponseWriter, *http.Request)
	PutProduct(http.ResponseWriter, *http.Request)
	DeleteProduct(http.ResponseWriter, *http.Request)
}
