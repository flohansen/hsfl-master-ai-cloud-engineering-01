package products

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/_mocks"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDefaultController(t *testing.T) {
	ctrl := gomock.NewController(t)

	productRepository := mocks.NewMockRepository(ctrl)
	controller := DefaultController{productRepository}

	t.Run("GetProducts", func(t *testing.T) {
		t.Run("should return 500 INTERNAL SERVER ERROR if query failed", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/products", nil)

			productRepository.
				EXPECT().
				FindAll().
				Return(nil, errors.New("query failed")).
				Times(1)

			// when
			controller.GetProducts(w, r)

			// then
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("should return all products", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/products", nil)

			productRepository.
				EXPECT().
				FindAll().
				Return([]*model.Product{{ID: 999}}, nil).
				Times(1)

			// when
			controller.GetProducts(w, r)

			// then
			res := w.Result()
			var response []model.Product
			err := json.NewDecoder(res.Body).Decode(&response)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.Len(t, response, 1)
			assert.Equal(t, int64(999), response[0].ID)
		})
	})

	t.Run("PostProducts", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if payload is not json", func(t *testing.T) {
			tests := []io.Reader{
				nil,
				strings.NewReader(`{"invalid`),
			}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/v1/products", test)

				// when
				controller.PostProducts(w, r)

				// then
				assert.Equal(t, http.StatusBadRequest, w.Code)
			}
		})

		t.Run("should return 400 BAD REQUEST if payload is incomplete", func(t *testing.T) {
			tests := []io.Reader{
				strings.NewReader(`{"price": 99.99}`),
				strings.NewReader(`{"description": "amazing product"}`),
				strings.NewReader(`{"retailer": "the company"}`),
			}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/v1/products", test)

				// when
				controller.PostProducts(w, r)

				// then
				assert.Equal(t, http.StatusBadRequest, w.Code)
			}
		})

		t.Run("should return 500 INTERNAL SERVER ERROR if persisting failed", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/api/v1/products",
				strings.NewReader(`{"name":"test product","retailer":"the company"}`))

			productRepository.
				EXPECT().
				Create([]*model.Product{{Name: "test product", Retailer: "the company"}}).
				Return(errors.New("database error"))

			// when
			controller.PostProducts(w, r)

			// then
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("should create new product", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/api/v1/products",
				strings.NewReader(`{"name":"test product","retailer":"the company"}`))

			productRepository.
				EXPECT().
				Create([]*model.Product{{Name: "test product", Retailer: "the company"}}).
				Return(nil)

			// when
			controller.PostProducts(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})

	t.Run("GetProduct", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if product id is not numerical", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/products/aaa", nil)
			r = r.WithContext(context.WithValue(r.Context(), "productid", "aaa"))

			// when
			controller.GetProduct(w, r)

			// then
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("should return 500 INTERNAL SERVER ERROR query failed", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/products/1", nil)
			r = r.WithContext(context.WithValue(r.Context(), "productid", "1"))

			productRepository.
				EXPECT().
				FindById(int64(1)).
				Return(nil, errors.New("database error"))

			// when
			controller.GetProduct(w, r)

			// then
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("should return 200 OK and product", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/v1/products/1", nil)
			r = r.WithContext(context.WithValue(r.Context(), "productid", "1"))

			productRepository.
				EXPECT().
				FindById(int64(1)).
				Return(&model.Product{ID: 1, Name: "test product", Retailer: "the company"}, nil)

			// when
			controller.GetProduct(w, r)

			// then
			res := w.Result()
			var response model.Product
			err := json.NewDecoder(res.Body).Decode(&response)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.Equal(t, int64(1), response.ID)
			assert.Equal(t, "test product", response.Name)
		})
	})

	t.Run("PutProduct", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if product id is not numerical", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/api/v1/products/aaa", nil)
			r = r.WithContext(context.WithValue(r.Context(), "productid", "aaa"))

			// when
			controller.PutProduct(w, r)

			// then
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("should return 400 BAD REQUEST if payload is not json", func(t *testing.T) {
			tests := []io.Reader{
				nil,
				strings.NewReader(`{"invalid`),
			}

			for _, test := range tests {
				// given
				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/api/v1/products/1", test)
				r = r.WithContext(context.WithValue(r.Context(), "productid", "1"))

				// when
				controller.PutProduct(w, r)

				// then
				assert.Equal(t, http.StatusBadRequest, w.Code)
			}
		})

		t.Run("should return 500 INTERNAL SERVER ERROR if query failed", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/api/v1/products/1",
				strings.NewReader(`{"id": 999}`))
			r = r.WithContext(context.WithValue(r.Context(), "productid", "1"))

			productRepository.
				EXPECT().
				Create([]*model.Product{{ID: 1}}).
				Return(errors.New("database error"))

			// when
			controller.PutProduct(w, r)

			// then
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("should update one product", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/api/v1/products/1",
				strings.NewReader(`{"id": 999}`))
			r = r.WithContext(context.WithValue(r.Context(), "productid", "1"))

			productRepository.
				EXPECT().
				Create([]*model.Product{{ID: 1}}).
				Return(nil)

			// when
			controller.PutProduct(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		t.Run("should return 400 BAD REQUEST if product id is not numerical", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/api/v1/products/aaa", nil)
			r = r.WithContext(context.WithValue(r.Context(), "productid", "aaa"))

			// when
			controller.DeleteProduct(w, r)

			// then
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("should return 500 INTERNAL SERVER ERROR if query fails", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/api/v1/products/1", nil)
			r = r.WithContext(context.WithValue(r.Context(), "productid", "1"))

			productRepository.
				EXPECT().
				Delete([]*model.Product{{ID: 1}}).
				Return(errors.New("database error"))

			// when
			controller.DeleteProduct(w, r)

			// then
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("should return 200 OK", func(t *testing.T) {
			// given
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/api/v1/products/1", nil)
			r = r.WithContext(context.WithValue(r.Context(), "productid", "1"))

			productRepository.
				EXPECT().
				Delete([]*model.Product{{ID: 1}}).
				Return(nil)

			// when
			controller.DeleteProduct(w, r)

			// then
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})
}
