package main

import "fmt"

func main() {

	fmt.Println("This is the main function in ch3")

	//appending to slices
	myslice := []int{10, 20, 30}
	myslice = append(myslice, 40)

	fmt.Println(myslice)

	//to check if slice is nil

	var ms []int
	fmt.Println(ms == nil)
	fmt.Println(ms)

	ms = append(ms, 10)
	fmt.Println(ms == nil)
	fmt.Println(ms)

}
