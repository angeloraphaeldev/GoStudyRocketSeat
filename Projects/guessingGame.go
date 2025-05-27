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
	fmt.Println("Welcome to Guessing Game!! 😎")
	fmt.Println(
		"\n ⚠️  A Random number will be choose. Try hit. The number is between 0 and 100")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)
	hits := [10]int64{}

	for i := range hits {
		fmt.Println("\nWhat your guess? 👀")
		scanner.Scan()
		hit := scanner.Text()
		hit = strings.TrimSpace(hit)

		hitInt, err := strconv.ParseInt(hit, 10, 64)
		if err != nil {
			fmt.Println("Use a integer number")
			return
		}

		switch {
		case hitInt < x:
			fmt.Println("🚫 Wrong Hit, X is higher than ", hitInt)
		case hitInt > x:
			fmt.Println("🚫 Wrong Hit, X is lower than ", hitInt)
		case hitInt == x:
			fmt.Printf(" 🎉🎉🎉 CONGRATULATIONS! You are RIGHT! The number was %d\n"+
				"You hit in %d tries.\n"+
				"That was yours hits: %v\n",
				x, i+1, hits[:i], //Catch all array values until i index
			)
			return
		}

		hits[i] = hitInt
	}
	fmt.Printf("\n😥 OH nooooo, you dont guess the number. The number was %d\n"+
		"Nice try, you should try again. That was your guesses%v\n", x, hits)
}
