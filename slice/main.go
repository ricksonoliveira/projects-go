package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	makeSlice := make([]int, 2)

	makeSlice[0] = 10
	makeSlice[1] = 20

	mkSlice := makeSlice[1:]
	mkSlice1 := makeSlice[:]
	mkSlice2 := makeSlice[1:2]
	makeSlice[1] = 200

	fmt.Println(makeSlice)
	fmt.Println(mkSlice)
	fmt.Println(mkSlice1)
	fmt.Println(mkSlice2)
}
