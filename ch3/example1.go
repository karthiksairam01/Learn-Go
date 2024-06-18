/*My first running piece of golang code!*/

package main

import "fmt"

func main() {

	myslice := []int{10, 20, 30}
	myslice = append(myslice, 40)

	fmt.Println(myslice)

}
