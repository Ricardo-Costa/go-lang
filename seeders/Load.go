package seeders

import "github.com/Ricardo-Costa/go-lang/entity"

var Products entity.Products

func LoadData() {
	p1 := entity.NewProduct()
	p1.Name = "Product #1"

	p2 := entity.NewProduct()
	p2.Name = "Product #2"

	Products.Add(*p1)
	Products.Add(*p2)
}
