package main

import "fmt"
var name string 

func main (){
	fmt.Println("what is your name")
	fmt.Scan(&name)
	greetings(name)
}


func greetings(name string){
	fmt.Printf("hello %s, nice to meet you",name)
}