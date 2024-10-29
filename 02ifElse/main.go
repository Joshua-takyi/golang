package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to the course for if and else ")
	correctPassword := 5
	loginCount := 3

	for {
		fmt.Println("enter the password")
		userInput := bufio.NewReader(os.Stdin)

		input, err := userInput.ReadString('\n')

		if err != nil {
			fmt.Println("an error occurred", err.Error())
		}
		// note: strconv (string convert ) Atoi is used to convert string to int
		convertedInput, convErr := strconv.Atoi(strings.TrimSpace(input))

		if convErr != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if convertedInput == correctPassword {
			fmt.Println("you have successfully logged in")
			break
		} else {
			fmt.Println("you have entered the wrong password")
			loginCount--
			fmt.Printf("you have %v trials left", loginCount)
		}

	}
	// note : inline if statement
	if num := 5; num < 2 {
		fmt.Println("number is less than 2")
	} else if num == 5 {
		fmt.Println("number is 5")
	} else {
		fmt.Println("fuck num")
	}

}
