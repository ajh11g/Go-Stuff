// Simple Number Guessing Game
// Description: system generates a random number.
//				user attempts to guess that number.
// Version: 1.0

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	startNum, endNum, numTries := 1, 20, 0
	targetNumber := random(startNum, endNum)
	var playerGuess int

	fmt.Printf("I'm thinking of a number between %d and %d. You have 3 guesses. ", startNum, endNum)

	for playerGuess != targetNumber {
		fmt.Println("Enter your guess: ")
		fmt.Scan(&playerGuess)

		if playerGuess == targetNumber {
			fmt.Println("CORRECT")
		} else if numTries == 2 {
			fmt.Println("You've reached the max number of tries. GAME OVER")
			break
		} else {
			fmt.Println("WRONG. Try Again!")
			numTries++
		}
	}
}
