package service

import (
	"database_implementation/internal"
)

type ProductDefault struct {
	rp internal.ProductRepository
}

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

func (s *ProductDefault) GetOne(id int) (internal.Product, error) {
	product, err := s.rp.GetOne(id)
	if err != nil {
		return internal.Product{}, err
	}
	return product, nil
}

func (s *ProductDefault) GetAll() ([]internal.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductDefault) Store(p *internal.Product) error {
	//TODO implement me
	panic("implement me")
}

func (s *ProductDefault) Update(p *internal.Product) (internal.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductDefault) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
