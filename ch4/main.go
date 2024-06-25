package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shadow() {

	x := 10

	if x > 5 {

		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}

	fmt.Println(x)
}

func shadow2() {

	var x int = 10

	if x > 5 {

		fmt.Println(x)
		var x int = 5
		fmt.Println(x)

	}

	fmt.Println(x)

}

func infinite() {

	for {

		fmt.Println("Hello")
	}
}

func labels() {

	samples := []string{"hello", "world"}

outer:

	for _, sample := range samples {

		for i, r := range sample {

			fmt.Println(i, r, string(r))

			if r == 'l' {

				continue outer
			}
		}

		fmt.Println()
	}
}

func q2() {

	n := rand.Intn(100)

	if n <= 33 {

		fmt.Println("Low: ", n)

	} else if n > 33 && n < 66 {

		fmt.Println("Mid: ", n)

	} else {

		fmt.Println("High: ", n)
	}
}

func fib() {

	a, b := 0, 1

	for i := 0; i < 10; i++ {

		fmt.Println(a)

		a, b = b, a+b
	}
}

func q3() {

	m := map[string]string{

		"Ally":   "S",
		"Bonnie": "A",
		"Kathy":  "C",
	}

	for k, v := range m {

		fmt.Printf("%s scored a %s grade\n", k, v)
	}
}

func q4() {

	today := time.Now().Weekday()

	switch today {

	case time.Monday:
		fmt.Println("It's Monday, start of the work week!")

	case time.Friday:
		fmt.Println("It's Friday, almost weekend!")

	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")

	default:
		fmt.Println("It's a regular workday")

	}

}

func main() {

	// fmt.Println("With shadow:")
	// shadow()

	// fmt.Println("With shadow2:")
	// shadow2()

	// infinite()

	//labels()

	// testing through questions

	// Write a Go program that uses an if-else if-else structure to categorize a randomly generated number as "low", "medium", or "high". Use the math/rand package to generate the number.

	q2()

	// Write a Go program that uses a complete style for loop to print the first 10 Fibonacci numbers.

	fib()

	// Write a Go program that uses a for-range loop to iterate over a map of student names and their grades, and print each student's name and grade.

	q3()

	// Write a Go program that uses a switch statement to print a message based on the day of the week. Use the time package to get the current day.

	q4()

}
