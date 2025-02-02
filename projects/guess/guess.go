package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("-----Welcome To The Guessing Game-----")
	fmt.Println("You have only 3 tries to guess the correct number between 0 and 9")

	// Generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)
	var guess int // User's guess variable

	// Loop to get user's input
	for try := 1; try > 0; try++ {
		fmt.Printf("[*] TRIAL %d\n", try)
		fmt.Print("\tPlease enter your number: ")
		fmt.Scan(&guess) // User's input

		// Check if the user's guess is correct
		if guess < secretNumber {
			fmt.Printf("\tSorry, wrong guess ; number is too small\n ")
		} else if guess > secretNumber {
			fmt.Printf("\t Sorry, wrong guess ; number is too large\n ")
		} else {
			fmt.Printf("\t You win!\n")
			break
		}

		// Check if the user has exhausted all the trials
		if try == 3 {
			fmt.Printf("Game over!!\n ")
			fmt.Printf("The correct number is %d\n", secretNumber)
			break
		}

	}
}
