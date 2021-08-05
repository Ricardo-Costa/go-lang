package entity

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID   string
	Name string `json:"full_name"`
}

func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}

type Products struct {
	Product []Product
}

func (p *Products) Add(Product Product) {
	p.Product = append(p.Product, Product)
}