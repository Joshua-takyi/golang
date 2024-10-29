package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("how much is the pizza:")

	// * try catch !not available so we use the comma Ok || err err  

	input,err := reader.ReadString('\n')

	fmt.Println("thank you for buying from us ", input)
	if err != nil {
		fmt.Println("error occurred")
	}else{
		fmt.Println("no error")
	}
}