package products

import "net/http"

type DefaultController struct {
}

func (ctrl *DefaultController) GetProducts(http.ResponseWriter, *http.Request) {
}

func (ctrl *DefaultController) PostProducts(http.ResponseWriter, *http.Request) {
}

func (ctrl *DefaultController) GetProduct(http.ResponseWriter, *http.Request) {
}

func (ctrl *DefaultController) PutProduct(http.ResponseWriter, *http.Request) {
}

func (ctrl *DefaultController) DeleteProduct(http.ResponseWriter, *http.Request) {
}
