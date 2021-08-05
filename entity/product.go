package entity

type Product struct {
	ID   int
	Name string
}

type Products struct {
	Product []Product
}

func (p *Products) Add(Product Product) {
	p.Product = append(p.Product, Product)
}