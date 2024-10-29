package main

import "fmt"

func main() {
	fmt.Println("welcome to the array classes")

	var fruitList []int

	fmt.Print("this is the type \n",len(fruitList))


	vegetableList := [] string {"mango","pawpaw","apple"}

	fmt.Printf("this is the type : %T\n",vegetableList)

	fmt.Println("this is the length", len(vegetableList))
}