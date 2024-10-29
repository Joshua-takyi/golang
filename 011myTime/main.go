package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	* this format("01-02-2006 Monday 15:04:05") gives us the Date, Day and Time
	* whiles format("01-02-2006") only gives us the Date eg("28/10/2024") 
	* whiles format("Monday") only gives us the day eg("Tuesday")
	* whiles format("15:04:05") only gives us the Time of the day eg("09:23:40") 
	*/


	createdTime := time.Date(1999,time.August,23,22,22,10,0,time.UTC)
	fmt.Println(createdTime.Format("01-02-2006"))
	fmt.Println(createdTime.Format("Monday"))
	fmt.Println(createdTime.Format("15:04:05"))
}
