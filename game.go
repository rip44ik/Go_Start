// guess - игра, в которой игрок должен угадать случайное число
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("Загадано число от 1 до 100")
	fmt.Println("Угадайте его за 10 попыток. Готовы ?")

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Вы имеете", 10-guesses, "попыток.")

		fmt.Print("Введите свое число:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Упс,загаданое число МЕНЬШЕ.")
		} else if guess < target {
			fmt.Println("Упс,загаданое число БОЛЬШЕ.")
		} else {
			success = true
			fmt.Println("Вы угадали! за", guesses+1, "попыток")
			break
		}
	}
	if !success {
		fmt.Println("Извините вы не угадали загаданное число. Загаданное число - ", target)
	}
}
