package products

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
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
		t.Run("should create new product", func(t *testing.T) {
			// given
			// when
			// then
		})
	})

	t.Run("GetProduct", func(t *testing.T) {
		t.Run("should return one product", func(t *testing.T) {
			// given
			// when
			// then
		})
	})

	t.Run("PutProduct", func(t *testing.T) {
		t.Run("should update one product", func(t *testing.T) {
			// given
			// when
			// then
		})
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		t.Run("should delete one product", func(t *testing.T) {
			// given
			// when
			// then
		})
	})
}
