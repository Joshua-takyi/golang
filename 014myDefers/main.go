package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Defer in Golang")

	defer fmt.Print("hello \n")
	defer fmt.Print("world\n")
	defer fmt.Print("friend\n")
	fmt.Println("Bye Bye")

	/*
		? li-fo  last in first out
		! this was the output the last element in the block of code with the defer outputs first followed by the one on top of it ..so the top most one executes last


			*Welcome to Defer in Golang
			*Bye Bye
			*friend
			*world
			*hello
	*/
}
