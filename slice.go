package main

import "fmt"

func main() {


	//slice := make([]int, 0)
	slice := make([]int, 0, 5)
	fmt.Println(slice)

	slice = append(slice, 1, 25, 31, 245)

	fmt.Println(slice, len(slice))
	fmt.Println(slice, cap(slice))


	sl := make([]int, 2, 3)
	for i := 0; i < 20; i++ {
		sl = append(sl, i)
		fmt.Println(sl)
		fmt.Println("Len :", len(sl), " Cap:", cap(sl))
	}

	// mantendo com ponteiro
	// sl = slice
	sl3 := slice
	slice[0] = 995

	// Obs.: Se capacidade do slice for alterada internamente o apontamento tambem será afetado

	//fmt.Println(sl)
	fmt.Println(sl3)
	fmt.Println(slice)

}
