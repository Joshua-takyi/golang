package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our shop")
	fmt.Println("please rate our setup")
	reader := bufio.NewReader(os.Stdin)
	input,err :=reader.ReadString('\n')

	if err != nil || input == "" {
		fmt.Println("an error occurred",err.Error())
	}else{
		//todo don't forget to add the error handling to the strconv.parseFloat
		numConvert,_	 := strconv.ParseFloat(strings.TrimSpace(input),64)
		numConvert = numConvert + 1
		fmt.Printf("type of numConvert : %T\n",numConvert)
		fmt.Println(numConvert)
	}
}