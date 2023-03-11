package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) Product {
	return Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func (p *Product) Update(name string, price float64) {
	p.Name = name
	p.Price = price
}

func (p *Product) Validate() (msgs []string) {
	if p.Name == "" {
		msgs = append(msgs, "Name is required")
	}
	if p.Price <= 0 {
		msgs = append(msgs, "Price must be greater than zero")
	}
	return
}
