package main

import "fmt"

func main() {

	// * (*) is used to get the value of in the variable
	// ! whiles (&) is used to get address of the variable being stored in the memory


	// // * this means we have created a pointer
	// var name *string

	// // * this returns the address of the name
	// myName := &name


	name := "joshua"

	myName :=&name

	// * this returns the address of the name
	fmt.Println(myName)
	// * this returns the value being stored in the name
	fmt.Println(*myName)
}