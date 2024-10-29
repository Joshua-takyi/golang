package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	// hope()

	added, myMessage := proFunc(1, 2, 3, 4, 5)
	fmt.Printf("the result is : %v \n", added)
	fmt.Printf("the message is : %v \n", myMessage)
}

// ! anonymous function
// var hope = func() {
// 	fmt.Println("Hope you are doing well")
// }

// ! normal functions
// func add(x int, y int) int {
// 	return x + y
// }

// ! pro function
func proFunc(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "into the fires of battle unto the anvil of war"
}
