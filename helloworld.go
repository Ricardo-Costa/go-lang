package main

import (
	"errors"
	"fmt"
	"github.com/Ricardo-Costa/go-lang/entity"
)

func main() {
	fmt.Println("Hello World!!")

	res, err := sum(10, 15)
	fmt.Println(res)

	res2, err2 := sum(-10, 15)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res2, err2.Error())

	product := entity.Product{}
	product.ID = 123
	product.Name = "Shirt"

	products := entity.Products{}
	products.Add(product)

	fmt.Println(product)
	fmt.Println(products)
}

func sum(a int, b int) (int, error) {
	if (a < 0) {
		return 0, errors.New("A variable is less than zero.")
	}
	return a + b, nil
}
