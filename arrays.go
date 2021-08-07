package main

import "fmt"

func main() {

	var x [10]int

	x[9] = 123

	z := [5]int { 1, 23, 4, 54, 6 }

	fmt.Println(x, z)
}
