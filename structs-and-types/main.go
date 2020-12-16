package main

import "fmt"

type User struct {
	Name    string
	Age     Age
	Address string
	Live    bool
}

func (u *User) ToDie() {
	fmt.Println("F")

	u.Live = false
}

type Age int

func (a *Age) ToAge() {
	*a++
}

func main() {
	user := User{"Rick", 23, "Address", true}
	userUnordered := User{
		Address: "R. Tal",
		Name:    "Rick O Burro",
		Live:    true,
		Age:     23,
	}

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(userUnordered)
	fmt.Println("---------------------")

	var a Age

	a = 23
	user.Age.ToAge()
	fmt.Println(a)
	user.ToDie()
}
