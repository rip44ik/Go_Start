package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	amount, err := paintNeeded(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}
}
