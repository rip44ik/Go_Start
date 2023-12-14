package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("Загадано число от 1 до 100")
	fmt.Println("Угадайте за 10 попыток. Готовы ?")
	fmt.Println(target)
}
