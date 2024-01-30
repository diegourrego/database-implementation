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

func (s *ProductDefault) GetAll() (products []internal.Product, err error) {
	products, err = s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductDefault) Save(p *internal.Product) error {
	if err := s.rp.Save(p); err != nil {
		return err
	}
	return nil
}

func (s *ProductDefault) Update(p *internal.Product) (internal.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductDefault) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
