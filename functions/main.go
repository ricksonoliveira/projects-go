package main

import "fmt"

func main() {
	multipleReturn, msg := multipleReturn(30, 15)
	if msg != "" {
		fmt.Println(msg)
	}
	resultSub, _ := sub(10, 5)
	result := soma(10, 20)
	fmt.Println("Result is", result)
	fmt.Println("Result sub is", resultSub)
	fmt.Println("Result multiple return is", multipleReturn)
}

func soma(n1, n2 int) int {

	num1, num2 := 5, 10
	fmt.Println(num1 + num2)

	return n2 + n1
}

func sub(n1, n2 int) (int, int) {

	num1, num2 := 7, 10
	fmt.Println(num1 + num2)

	return 10, n1 - n2
}

func multipleReturn(n1, n2 int) (int, string) {
	fmt.Println("Inciando ...")
	result := n1 + n2
	if result > 10 {
		return result, "Value's higher than 10"
	}
	return result, ""
}
