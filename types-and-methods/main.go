package main

import "fmt"

type List []interface{}

func main() {
	var l List

	l.Init()

	l.Show()
}

func (l *List) Init() {
	*l = List{
		10,
		"Hello",
		1.5,
		false,
	}
}

func (l List) Show() {
	fmt.Println(l...)
}
