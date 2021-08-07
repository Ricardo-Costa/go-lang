package main

import "fmt"

func main() {


	//slice := make([]int, 0)
	slice := make([]int, 0, 5)
	fmt.Println(slice)

	slice = append(slice, 1, 25, 31, 245)

	fmt.Println(slice, len(slice))
	fmt.Println(slice, cap(slice))

}
