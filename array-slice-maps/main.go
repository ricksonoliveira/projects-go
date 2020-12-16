package main

import "fmt"

func main() {
	math := [10]int{}
	math1 := [10]int{0, 5, 10}
	math2 := [3]int{5}

	user := map[string]string{
		"name": "Rick O Burro",
		"nick": "rato",
	}

	user["age"] = "23"

	slice := []int{0, 5, 10}
	slice = append(slice, 90, 110, 45)

	fmt.Println(math)
	fmt.Println(math1)
	fmt.Println(math2)
	fmt.Println(math2)
	fmt.Println(user)
	fmt.Println(user["nick"])
	fmt.Println(user["age"])
	fmt.Println(slice)

}
