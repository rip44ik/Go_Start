package main

import (
	"fmt"
)

//var a, b, c int

func main() {
	message, entered := enterTheClub(24)
	fmt.Println(message)
	fmt.Println(entered)

}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "Входи, только акуратно", true
	} else {
		return "Тебе нет 18-ти", false
	}
}
