package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)+1

	fmt.Println("Let's play a guessing game. \nI've picked a number between 1 and 100, try to guess it.")

	for i := 1; i <= 5; i++ {
		fmt.Printf("This is try %d\n", i)

		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("only input numbers please")
			return
		}

		if guess < target {
			fmt.Println("Too low.")
		} else if guess > target {
			fmt.Println("Too high.")
		} else {
			fmt.Println("You're right!")
			return
		}
	}

	fmt.Printf("Sorry, you failed. It was %d\n", target)
}
