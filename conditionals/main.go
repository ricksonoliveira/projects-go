package main

import "fmt"

func main() {
	b := 4

	if b == 4 { //OPERADORES:  !== diferente de  &&  and  || ou

		fmt.Println("IF")

	} else {

		fmt.Println("ELSE")

	}

	if variable := 3; b > variable { //OPERADORES:  !== diferente de  &&  and  || ou

		fmt.Println("IF MAIOR")

	} else {

		fmt.Println("ELSE MENOR ")

	}

	num := 10

	switch num {
	case 2:
		fmt.Println("Number 2")
	case 3:
		fmt.Println("Number 3")
	default:
		fmt.Println("Default")
	}
}
