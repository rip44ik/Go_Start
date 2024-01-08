package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func main() {

	x := Contact{"Я", 33}
	x.name = "Евгений"
	fmt.Println(x.age)
	fmt.Println(x.name)
}
