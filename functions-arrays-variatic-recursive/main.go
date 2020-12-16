package main

import "fmt"

func main() {
	soma(12, 15, 40, 25, 80)
	fmt.Println(variaticaFunc(12, 15, 40, 25, 80))
}

func soma(nums ...int) { // variatic function
	fmt.Println(nums)
	fakeLoopRecur()
}

func variaticaFunc(nums ...int) int { // variatic function with for range
	var total int

	for _, num := range nums {
		total += num
	}

	return total
}

var i = 10

func fakeLoopRecur() { // recursive function
	fmt.Println(i)
	i--
	if i > 0 {
		fakeLoopRecur()
	}
}
