package main

import "fmt"

func main() {

	num := 20

	// while
	for num > 0 {
		fmt.Println(num)
		num--
	}

	for num := 0; num < 20; num++ {
		fmt.Println(num)
	}
}
