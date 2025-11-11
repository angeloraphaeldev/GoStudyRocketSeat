package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Guessing Game!! 😎")
	fmt.Println(
		"\n ⚠️  A Random number will be choose. Try hit. The number is between 0 and 100")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("\nWhat's your guess? 👀")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Please enter a valid integer number")
			continue
		}

		if guessInt < x {
			fmt.Println("🚫 Incorrect guess, the number is higher than ", guessInt)
		} else if guessInt > x {
			fmt.Println("🚫 Incorrect guess, the number is lower than ", guessInt)
		} else {
			fmt.Printf(" 🎉🎉🎉 CONGRATULATIONS! You guessed it RIGHT! The number was %d\n"+
				"You hit in %d attempts.\n"+
				"That was your guesses: %v\n",
				x, i+1, guesses[:i],
			)
			return
		}

		guesses[i] = guessInt
	}
	fmt.Printf("\n😥 OH NO, you didnt guess the number. The number was %d\n"+
		"Good try, you should try again. Your guesses where: %v\n", x, guesses)
}
