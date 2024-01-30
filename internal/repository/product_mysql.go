package repository

import (
	"database/sql"
	"database_implementation/internal"
	"errors"
	"github.com/go-sql-driver/mysql"
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

func (r *ProductMysql) GetAll() (products []internal.Product, err error) {
	rows, err := r.db.Query(
		"SELECT p.`id`, p.`name`, p.`type`, p.`count`, p.`price` FROM `products` AS p")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p internal.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return
}

func (r *ProductMysql) Save(p *internal.Product) (err error) {
	// Debemos ejecutar la query
	query := "INSERT INTO products (`name`, `type`, `count`, `price`) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		switch mysqlErr.Number {
		case 1062:
			err = internal.ErrProductDuplicated
		default:
			return err
		}
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	p.ID = int(id)
	return
}

func (r *ProductMysql) Update(p *internal.Product) (internal.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ProductMysql) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
