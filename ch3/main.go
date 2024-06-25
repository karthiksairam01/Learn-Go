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

	fmt.Println("------Testing question 1------")

	q1()

}

func q1() {

	type Book struct {
		Title  string
		Author string
		Pages  int
	}

	newbook := Book{

		Title:  "Learning Go",
		Author: "Jon Bodner",
		Pages:  500,
	}

	fmt.Printf("The title of the book is: %s whose author is %s having %d pages! \n", newbook.Title, newbook.Author, newbook.Pages)

	type Library struct {
		Books []Book
	}

	newlib := Library{}

	fmt.Printf("Library before adding books: %v \n", newlib)

	book1 := Book{"The Go Programming Language", "Alan A. A. Donovan", 380}
	book2 := Book{"Learning Go", "Jon Bodner", 300}

	newlib.Books = append(newlib.Books, book1)
	newlib.Books = append(newlib.Books, book2)

	fmt.Printf("Library after adding books: %v \n", newlib)

}
