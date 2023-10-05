package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products/model"
)

type createProductRequest struct {
	Name        string  `json:"name"`
	Retailer    string  `json:"retailer"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

type updateProductRequest struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Retailer    string  `json:"retailer"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

func (r createProductRequest) isValid() bool {
	return r.Name != "" && r.Retailer != ""
}

func (r updateProductRequest) isValid() bool {
	return r.ID != 0
}

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

func (ctrl *DefaultController) PostProducts(w http.ResponseWriter, r *http.Request) {
	var request createProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ctrl.productRepository.Create([]*model.Product{{
		Name:        request.Name,
		Retailer:    request.Retailer,
		Price:       request.Price,
		Description: request.Description,
	}}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ctrl *DefaultController) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.Context().Value("productid").(string)

	id, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := ctrl.productRepository.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (ctrl *DefaultController) PutProduct(w http.ResponseWriter, r *http.Request) {
	var request updateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ctrl.productRepository.Create([]*model.Product{{
		ID:          request.ID,
		Name:        request.Name,
		Retailer:    request.Retailer,
		Price:       request.Price,
		Description: request.Description,
	}}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ctrl *DefaultController) DeleteProduct(http.ResponseWriter, *http.Request) {
}
