package main

import "fmt"

func main() {
	fmt.Println("welcome to the course of Maps (key : value pairs)")
	
		// note : we have to use the map to tell the machine that we are using maps 
		
	names:=make(map[string]string)

	names["jo"]="joshua"
	names["ph"]="philip"
	names["tm"]="tom"
	fmt.Println(names)
	// note : we use the delete key to delete an element from the map delete[map,"key of the value"]
	delete(names,"jo")
	fmt.Println(names)

	// note : loop (for-each loop)
	/*
		? key,value := range map{
		? fmt.Printf("what ever you want to print",key,value)
		?}
	*/


	for key,value := range names{
		fmt.Printf("key is %v and value is %v\n",key,value)
	} 
}