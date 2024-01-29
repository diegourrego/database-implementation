package internal

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"quantity"`
	Price float64 `json:"price"`
}

type ProductRepository interface {
	GetOne(id int) (Product, error)
	GetAll() ([]Product, error)
	Store(p *Product) error
	Update(p *Product) (Product, error)
	Delete(id int) error
}

type ProductService interface {
	GetOne(id int) (Product, error)
	GetAll() ([]Product, error)
	Store(p *Product) error
	Update(p *Product) (Product, error)
	Delete(id int) error
}
