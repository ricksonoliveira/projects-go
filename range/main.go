package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4} //[...] means the array may have unlimited range

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("-----------------------------")

	for key, value := range arr {
		fmt.Println(key, value)
	}

	fmt.Println("-----------------------------")

	user := map[string]string{
		"name": "Rick",
		"nick": "rato",
	}

	for key, value := range user {
		fmt.Printf("The field \"%s\" has the value equal to \"%s\"\n", key, value)
	}

	fmt.Println("-----------------------------")

	langs := []string{
		"Go",
		"Js",
		"Rust",
		"Python",
		"Elixir",
		"Haskel",
		"PHP",
	}

	for _, value := range langs {
		fmt.Printf("The langs available at the market are \"%s\"\n", value)
	}
}
