package products

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/flohansen/hsfl-master-ai-cloud-engineering/lib/database"
	"github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products/model"
)

type PsqlRepository struct {
	db *sql.DB
}

func NewPsqlRepository(config database.Config) (*PsqlRepository, error) {
	db, err := sql.Open("postgres", config.Dsn())
	if err != nil {
		return nil, err
	}

	return &PsqlRepository{db}, nil
}

const createProductsTable = `
create table if not exists products (
	id          serial  primary key,
	name        text    not null,
	retailer    text    not null,
	price       decimal not null default 0,
	description text             default ''
)
`

func (repo *PsqlRepository) Migrate() error {
	_, err := repo.db.Exec(createProductsTable)
	return err
}

const createProductsBatchQuery = `
insert into products (name, retailer, price, description) values %s
`

func (repo *PsqlRepository) Create(products []*model.Product) error {
	placeholders := make([]string, len(products))
	values := make([]interface{}, len(products)*4)

	for i := 0; i < len(products); i++ {
		placeholders[i] = fmt.Sprintf("($%d,$%d,$%d,$%d)", i*4+1, i*4+2, i*4+3, i*4+4)
		values[i*4+0] = products[i].Name
		values[i*4+1] = products[i].Retailer
		values[i*4+2] = products[i].Price
		values[i*4+3] = products[i].Description
	}

	query := fmt.Sprintf(createProductsBatchQuery, strings.Join(placeholders, ","))
	_, err := repo.db.Exec(query, values...)
	return err
}

const findAllProductsQuery = `
select id, name, retailer, price, description from products
`

func (repo *PsqlRepository) FindAll() ([]*model.Product, error) {
	rows, err := repo.db.Query(findAllProductsQuery)
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Retailer, &product.Price, &product.Description); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

const findProductByIdQuery = `
select id, name, retailer, price, description from products where id = $1 limit 1
`

func (repo *PsqlRepository) FindById(id int64) (*model.Product, error) {
	row := repo.db.QueryRow(findProductByIdQuery, id)

	var product model.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Retailer, &product.Price, &product.Description); err != nil {
		return nil, err
	}

	return &product, nil
}

const deleteProductsByIdQuery = `
delete from products where id in (%s)
`

func (repo *PsqlRepository) Delete(products []*model.Product) error {
	placeholders := make([]string, len(products))
	ids := make([]interface{}, len(products))

	for i := 0; i < len(products); i++ {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		ids[i] = products[i].ID
	}

	query := fmt.Sprintf(deleteProductsByIdQuery, strings.Join(placeholders, ","))
	_, err := repo.db.Exec(query, ids...)
	return err
}
