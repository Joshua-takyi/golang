package main

import (
	"fmt"
)

func main() {

	// * commented to make the code cleaner
	// fruitList := []string{"mango", "pawpaw", "pineapple"}

	// for {
		// fmt.Println("please enter the fruit you want to add in the list")
		// addedData := bufio.NewReader(os.Stdin)

		// input, err := addedData.ReadString('\n')
		// if err != nil {
		// 	fmt.Println("an error occurred while reading input", err.Error())
		// }
		// input = strings.TrimSpace(input)
		// fruitList = append(fruitList, input)
		// fmt.Println(fruitList)

		// // note [x:y]  the index of x is included and the index of y is excluded
		// if len(fruitList) == 5 {
		// 	fmt.Println("now we take ours", fruitList[2:4])
		// }
	// }



	// ! MAKE


	highSchoolNames := make([]string,4)

	highSchoolNames[0]="joshua"
	highSchoolNames[1]="philip"
	highSchoolNames[2]="thomas"
	highSchoolNames[3]="peter"
	/*
	note: since the make([]type,range) is 4 in our highSchoolNames after adding the index of 4 we got this error
	! panic: runtime error: index out of range [4] with length 4
	*/
	//? commented out to make the code cleaner
	// highSchoolNames[4]="patrick"s


	// * but when we use the append method is reallocates the memory and all the data will be reallocated

	// note: we wont get an error because we reallocated the memory 
	// ? this was the output [joshua philip thomas peter patrick]
	highSchoolNames = append(highSchoolNames, "patrick")
	fmt.Println(highSchoolNames)
	// fmt.Println(sort.StringsAreSorted(highSchoolNames))

	// note: (sort) an inbuilt function used for sorting of slices

	// sort.Strings(highSchoolNames)
	// fmt.Println(sort.StringsAreSorted(highSchoolNames))


	// note: removing a value from a slice based on its index
	/* 
	! FORMULA
	* dataToBeRemoved := index
	note : since the element at the end of the [:y] is excluded, the  value being stored in that index is removed and  added to the [x:] +1 which means add all the data which is after the deleted index
	* slice = append (slice[:dataToBeRemoved], slice[dataToBeRemoved +1 :]...) 
	*/
	dataToBeRemoved :=2

	
    highSchoolNames = append(highSchoolNames[:dataToBeRemoved], highSchoolNames[dataToBeRemoved+1:]...)
    fmt.Println(highSchoolNames) // Output: [joshua philip peter patrick]
}


