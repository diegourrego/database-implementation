package repository

import (
	"database/sql"
	"database_implementation/internal"
	"errors"
)

type ProductMysql struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductMysql {
	return &ProductMysql{db}
}

func (r *ProductMysql) GetOne(id int) (internal.Product, error) {
	row := r.db.QueryRow("SELECT p.`id`, p.`name`, p.`type`, p.`count`, p.`price` FROM `products` AS p WHERE p.id=?", id)
	if err := row.Err(); err != nil {
		return internal.Product{}, err
	}
	var product internal.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.Product{}, internal.ErrProductNotFound
		}
		return internal.Product{}, err
	}
	return product, nil
}

func (r *ProductMysql) GetAll() ([]internal.Product, error) {
	_, err := r.db.Query("SELECT p.`id`, p.`name`, p.`type`, p.`count`, p.`price` FROM `products` AS p")
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (r *ProductMysql) Store(p *internal.Product) error {
	//TODO implement me
	panic("implement me")
}

func (r *ProductMysql) Update(p *internal.Product) (internal.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ProductMysql) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
