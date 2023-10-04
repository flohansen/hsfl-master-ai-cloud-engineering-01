package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRouter(t *testing.T) {
	ctrl := gomock.NewController(t)

	productsHandler := mocks.NewMockHandler(ctrl)
	router := New(productsHandler)

	t.Run("should return 404 NOT FOUND if path is unknown", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/unknown/route", nil)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("should call products handler", func(t *testing.T) {
		// given
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/api/v1/products", nil)

		productsHandler.
			EXPECT().
			ServeHTTP(w, r).
			Times(1)

		// when
		router.ServeHTTP(w, r)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
