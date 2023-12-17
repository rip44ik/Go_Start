package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())
	x := 3
	switch x {
	case 2:
		x += 1
	case 3:
		x += 1
	case 4:
		x -= 1
	}
	fmt.Println(x)
}
