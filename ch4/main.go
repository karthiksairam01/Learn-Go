package main

import (
	"fmt"
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

func main() {

	fmt.Println("With shadow:")
	shadow()

	fmt.Println("With shadow2:")
	shadow2()

}
