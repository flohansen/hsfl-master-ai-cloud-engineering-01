package products

import (
	"encoding/json"
	"net/http"
)

type DefaultController struct {
	productRepository Repository
}

func NewDefaultController(
	productRepository Repository,
) *DefaultController {
	return &DefaultController{productRepository}
}

func (ctrl *DefaultController) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ctrl.productRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (ctrl *DefaultController) PostProducts(http.ResponseWriter, *http.Request) {
}

func (ctrl *DefaultController) GetProduct(http.ResponseWriter, *http.Request) {
}

func (ctrl *DefaultController) PutProduct(http.ResponseWriter, *http.Request) {
}

func (ctrl *DefaultController) DeleteProduct(http.ResponseWriter, *http.Request) {
}
