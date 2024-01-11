package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ { // fdsfdsfsd
		PrintNumber(i) //fdsfsdf
	}
}

//revive:disable:exported

func PrintHello() {
	fmt.Println("Hello, Go")
}

// PrintNumber this function using the fmt.Println function
//
//revive:enable:exported
func PrintNumber(number int) {
	fmt.Println(number)
}
