package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to my Switch course")
	// note Create a new random source and generator once outside the loop
	/*
	*we use the rand which comes from Math/rand to generate a random
	*it first used to be rand.seed but it got deprecated and now we use rand.NewSource
	 */
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	for {
		//? Generate a random number between 1 and 6
		generatedRan := random.Intn(6) + 1

		switch generatedRan {
		case 1:
			fmt.Println("You rolled a 1! Try again.")
		case 2:
			fmt.Println("You got a 2. Almost there!")
		case 3:
			fmt.Println("3! Halfway through.")
		case 4:
			fmt.Println("That's a 4. Nice roll!")
		case 5:
			fmt.Println("5! So close to the top.")
			fallthrough
		case 6:
			fmt.Println("6! You hit the jackpot!")
			fallthrough
		default:
			fmt.Println("Error: Value out of range")
		}
		//note: setTimeOut (1000)
		time.Sleep(time.Second)
	}
}
