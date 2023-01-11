// Game to guess a random number between 1 to 100

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("\nWelcome To The Guess Game \n")
	fmt.Println("You have to guess a number between 1 to 100")
	fmt.Println("You'll get 7 chances and a hint after every chance \n")

	fmt.Println("If you'll win I'll tell you a programming joke \n")

	// fmt.Println(target) // Uncomment this line to cheat!

	reader := bufio.NewReader(os.Stdin)
	success := false

	for guesses := 7; guesses > 0; guesses-- {
		fmt.Println("Gueses Remaining:", guesses)
		fmt.Print("Make a guess: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops! Your guess was low")
		} else if guess > target {
			fmt.Println("Opps! Your guess was high")
		} else {
			fmt.Println("That's correct! You won!")
			fmt.Println("Here's the Joke:")
			fmt.Println("There are only 10 kinds of people in this world: those who know binary and those who don’t.")
			success = true
			break
		}
		fmt.Println("")

	}
	if !success {
		fmt.Println("Sorry you lost. The correct answer was:", target)
	}
	for true {

	}
}
