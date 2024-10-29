package main

import "fmt"

type Add struct {
	Width, Length float64
}

func main() {
	fmt.Println("Welcome to Methods in Golang")

	rec := Add{Width: 10, Length: 10}
	fmt.Println(rec.Area())
}

func (a Add) Area() float64 {
	return a.Width * a.Length
}
