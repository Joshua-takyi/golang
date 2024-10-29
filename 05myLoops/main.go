package main

import "fmt"

func main() {
	fmt.Println("welcome to  loops in Golang")

	// days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	// fmt.Println(days)
	// ! remember when using i instead of i,e it returns only the index

	// note loops through the entire array
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }
	// note:loops through an array or slice
	// ! so basically only the index is being returned
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// note : i,e returns both the index and the value
	// ? : this is basically a foreach loop

	// for i, e := range days {
	// 	fmt.Printf("index is %v and value is %v\n", i, e)
	// }

	// note : this is somehow a while loop

	forNum := 1

	for forNum <= 10 {

		// ! break in loop

		// if forNum == 5 {
		// 	fmt.Println("the number reached it's mark")
		// 	break
		/*
			note: result
			value of forNum is:  1
			value of forNum is:  2
			value of forNum is:  3
			value of forNum is:  4
			the number reached it's mark

		*/
		// }

		// ! using the goto statement here
		if forNum == 3 {
			goto loc
		}
		// if forNum == 5 {
		// note : we add the forNum++ for it to iterates over the removed element
		// forNum++
		// fmt.Println("the number reached it's mark")
		// continue

		/*
			note: results
				welcome to  loops in Golang
				value of forNum is:  1
				value of forNum is:  2
				value of forNum is:  3
				value of forNum is:  4
				value of forNum is:  6
				value of forNum is:  7
				value of forNum is:  8
				value of forNum is:  9
				value of forNum is:  10


		*/
		// }

		// fmt.Println("value of forNum is: ", forNum)
		// forNum++
	}
	// note :goto statement
loc:
	fmt.Println("welcome to the location")

}
