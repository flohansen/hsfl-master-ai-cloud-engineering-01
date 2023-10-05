package products

import (
	"context"
	"database/sql"
	"testing"

	"github.com/flohansen/hsfl-master-ai-cloud-engineering/lib/containerhelpers"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/lib/database"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products/model"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestIntegrationPsqlRepository(t *testing.T) {
	postgres, err := containerhelpers.StartPostgres()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		postgres.Terminate(context.Background())
	})

	port, err := postgres.MappedPort(context.Background(), "5432")
	if err != nil {
		t.Fatalf("could not get database container port: %s", err.Error())
	}

	repository, err := NewPsqlRepository(database.PsqlConfig{
		Host:     "localhost",
		Port:     port.Int(),
		Username: "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	if err != nil {
		t.Fatalf("could not create products repository: %s", err.Error())
	}
	t.Cleanup(clearTables(t, repository.db))

	t.Run("Migrate", func(t *testing.T) {
		t.Run("should create products table", func(t *testing.T) {
			t.Cleanup(clearTables(t, repository.db))

			// given
			// when
			err := repository.Migrate()

			// then
			assert.NoError(t, err)
			assertTableExists(t, repository.db, "products", []string{"id", "name", "retailer", "price", "description"})
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("should create products", func(t *testing.T) {
			t.Cleanup(clearTables(t, repository.db))

			// given
			products := []*model.Product{
				{
					Name:     "test product 1",
					Retailer: "the company",
				},
				{
					Name:     "test product 2",
					Retailer: "the company",
				},
			}

			// when
			err := repository.Create(products)

			// then
			assert.NoError(t, err)
			assert.NotNil(t, getProductFromDatabase(t, repository.db, "test product 1"))
			assert.NotNil(t, getProductFromDatabase(t, repository.db, "test product 2"))
		})
	})

	t.Run("FindAll", func(t *testing.T) {
		t.Run("should return all products", func(t *testing.T) {
			t.Cleanup(clearTables(t, repository.db))

			// given
			products := []*model.Product{
				{
					Name:     "test product 1",
					Retailer: "the company",
				},
				{
					Name:     "test product 2",
					Retailer: "the company",
				},
			}

			for _, product := range products {
				insertProduct(t, repository.db, product)
			}

			// when
			products, err := repository.FindAll()

			// then
			assert.NoError(t, err)
			assert.Len(t, products, 2)
		})
	})

	t.Run("FindById", func(t *testing.T) {
		t.Run("should return one product", func(t *testing.T) {
			t.Cleanup(clearTables(t, repository.db))

			// given
			products := []*model.Product{
				{
					Name:     "test product 1",
					Retailer: "the company",
				},
				{
					Name:     "test product 2",
					Retailer: "the company",
				},
			}

			for _, product := range products {
				insertProduct(t, repository.db, product)
			}

			// when
			id := getProductFromDatabase(t, repository.db, "test product 1").ID
			product, err := repository.FindById(id)

			// then
			assert.NoError(t, err)
			assert.NotNil(t, product)
			assert.Equal(t, "test product 1", product.Name)
			assert.Equal(t, "the company", product.Retailer)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete products", func(t *testing.T) {
			t.Cleanup(clearTables(t, repository.db))

			// given
			products := []*model.Product{
				{
					Name:     "test product 1",
					Retailer: "the company",
				},
				{
					Name:     "test product 2",
					Retailer: "the company",
				},
			}

			for _, product := range products {
				insertProduct(t, repository.db, product)
				product.ID = getProductFromDatabase(t, repository.db, product.Name).ID
			}

			// when
			err := repository.Delete([]*model.Product{products[1]})

			// then
			assert.NoError(t, err)
			assert.NotNil(t, getProductFromDatabase(t, repository.db, "test product 1"))
			assert.Nil(t, getProductFromDatabase(t, repository.db, "test product 2"))
		})
	})
}

func getProductFromDatabase(t *testing.T, db *sql.DB, name string) *model.Product {
	row := db.QueryRow(`select id from products where name = $1`, name)

	var product model.Product
	if err := row.Scan(&product.ID); err != nil {
		return nil
	}

	return &product
}

func insertProduct(t *testing.T, db *sql.DB, user *model.Product) {
	_, err := db.Exec(`insert into products (name, retailer) values ($1, $2)`, user.Name, user.Retailer)
	if err != nil {
		t.Logf("could not insert product: %s", err.Error())
		t.FailNow()
	}
}

func clearTables(t *testing.T, db *sql.DB) func() {
	return func() {
		if _, err := db.Exec("delete from products"); err != nil {
			t.Logf("could not delete rows from products: %s", err.Error())
			t.FailNow()
		}
	}
}

func assertTableExists(t *testing.T, db *sql.DB, name string, columns []string) {
	rows, err := db.Query(`select column_name from information_schema.columns where table_name = $1`, name)
	if err != nil {
		t.Fail()
		return
	}

	scannedCols := make(map[string]struct{})
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			t.Logf("expected")
			t.FailNow()
		}

		scannedCols[column] = struct{}{}
	}

	if len(scannedCols) == 0 {
		t.Logf("expected table '%s' to exist, but not found", name)
		t.FailNow()
	}

	for _, col := range columns {
		if _, ok := scannedCols[col]; !ok {
			t.Logf("expected table '%s' to have column '%s'", name, col)
			t.Fail()
		}
	}
}
