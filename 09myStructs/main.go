package main

import "fmt"

// note : since most of the data will be exported or used by other packages it is recommended to start them with capital casing

	type Student struct{
		Name string
		Age int
		Grade float32
		Email string
		IsGrown bool
	}

func main() {
	fmt.Println("welcome to the course of structs")

	josh :=Student{
		"joshua",
		22,
		2.8,
		"william.henry.harrison@example-pet-store.com",
		false,
	}
	fmt.Println(josh)
	// ? we use the %+v to get more details
	fmt.Printf("data: %+v\n",josh)
	

	// note: if we want a single element from the Structs we use the data.key

	fmt.Printf("this is the name: %v \n and this is the email: %v \n", josh.Name,josh.Email)
}