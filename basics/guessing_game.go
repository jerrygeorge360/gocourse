package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1
	fmt.Println("welcome to the guessing game")
	fmt.Println("I have chosen a numbrer between 1 and 100")
	fmt.Println("can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("congratulations you guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("low")
		} else {
			fmt.Println("high")
		}
	}

}
