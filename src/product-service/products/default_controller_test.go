package products

import "testing"

func TestDefaultController(t *testing.T) {
	_ = DefaultController{}

	t.Run("GetProducts", func(t *testing.T) {
		t.Run("should return all products", func(t *testing.T) {
			// given
			// when
			// then
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
