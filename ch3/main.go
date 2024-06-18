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

	//len

	z := len(ms)
	fmt.Println("z is:", z)

	//append one slice to another

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6, 3}
	s3 := append(s1, s2...)
	fmt.Println(s1, s2, s3)

	//check capacity

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	//make

	p := make([]int, 10)
	fmt.Println(len(p), cap(p))

	q := make([]int, 5, 10)
	q = append(q, 5, 4, 6)
	fmt.Println(q, len(q), cap(q))

}
