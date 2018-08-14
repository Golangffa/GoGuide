//Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

//Run code
package main

import "fmt"

//The (int, int) in this function signature shows that the function returns 2 ints.

func vals() (int, int, int) {
	return 3, 7, 5
}
func main() {
	//Here we use the 2 different return values from the call with multiple assignment.

	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//If you only want a subset of the returned values, use the blank identifier _.

	_, _, d := vals()
	fmt.Println(d)
}

//$ go run multiple-return-values.go
//3
//7
//7
