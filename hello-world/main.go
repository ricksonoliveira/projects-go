package main

import "fmt"

var minhaVar = "nao pode ponto igual"

func main() {

	var varLocal = "Minha var local"

	name, age, address := "Rick Burro", 23, "R. Tal Apto 302 BL 05"
	num1, num2 := 40, 10

	var (
		name2    = "Meu nome 2"
		address2 = "Rua Fulano"
	)

	fmt.Print(name, age, address, minhaVar)
	fmt.Print(num1 * num2)
	fmt.Print(num1 / num2)
}
