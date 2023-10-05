package products

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products/model"
	"github.com/stretchr/testify/assert"
)

func TestPsqlRepository(t *testing.T) {
	db, dbmock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	repository := PsqlRepository{db}

	t.Run("Create", func(t *testing.T) {
		t.Run("should insert products in batches", func(t *testing.T) {
			// given
			products := []*model.Product{
				{
					Name:     "test product 1",
					Retailer: "test company",
				},
				{
					Name:     "test product 2",
					Retailer: "test company",
				},
			}

			dbmock.ExpectExec(`insert into products \(name, retailer, price, description\) values \(\$1,\$2,\$3,\$4\),\(\$5,\$6,\$7,\$8\)`).
				WithArgs("test product 1", "test company", sqlmock.AnyArg(), sqlmock.AnyArg(), "test product 2", "test company", sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(0, 2))

			// when
			err := repository.Create(products)

			// then
			assert.NoError(t, err)
			assert.NoError(t, dbmock.ExpectationsWereMet())
		})
	})

	t.Run("FindAll", func(t *testing.T) {
		t.Run("should return all products", func(t *testing.T) {
			// given
			dbmock.ExpectQuery(`select (.*) from products`).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "retailer", "price", "description"}).
					AddRow(1, "test product 1", "the company", 99.99, "description").
					AddRow(2, "test product 2", "the company", 9.99, "description"))

			// when
			products, err := repository.FindAll()

			// then
			assert.NoError(t, err)
			assert.NoError(t, dbmock.ExpectationsWereMet())
			assert.Len(t, products, 2)
			assert.Equal(t, "test product 1", products[0].Name)
			assert.Equal(t, "test product 2", products[1].Name)
		})
	})

	t.Run("FindById", func(t *testing.T) {
		t.Run("should return product by id", func(t *testing.T) {
			// given
			var id int64 = 999

			dbmock.ExpectQuery(`select (.*) from products where id = \$1 limit 1`).
				WithArgs(999).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "retailer", "price", "description"}).
					AddRow(1, "test product 1", "the company", 99.99, "description"))

			// when
			product, err := repository.FindById(id)

			// then
			assert.NoError(t, err)
			assert.NotNil(t, product)
			assert.NoError(t, dbmock.ExpectationsWereMet())
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete products in batch", func(t *testing.T) {
			// given
			products := []*model.Product{
				{
					ID:       1,
					Name:     "test product 1",
					Retailer: "test company",
				},
				{
					ID:       2,
					Name:     "test product 2",
					Retailer: "test company",
				},
			}

			dbmock.ExpectExec(`delete from products where id in \(\$1,\$2\)`).
				WithArgs(1, 2).
				WillReturnResult(sqlmock.NewResult(0, 2))

			// when
			err := repository.Delete(products)

			// then
			assert.NoError(t, err)
			assert.NoError(t, dbmock.ExpectationsWereMet())
		})
	})
}
