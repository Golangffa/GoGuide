package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Add performs addition of two numbers
func Add(x, y int) int {
	return x + y
}

// Subtract performs subtraction of two numbers
func Subtract(x, y int) int {
	return x - y
}

// Multiply performs multiplication of two numbers
func Multiply(x, y int) int {
	return x * y
}

// Divide performs division of two numbers
func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}
	return x / y, nil
}

func main() {

// s for subtraction
// p for product
// d for division
// m for multiplication

        arg1, _ := strconv.Atoi(os.Args[1])
        arg2, _ := strconv.Atoi(os.Args[2])
        arg3 := os.Args[3]

switch arg3 {
	case "s":
		fmt.Println(Subtract(arg1, arg2))
	case "p":
		fmt.Println(Add(arg1, arg2))
	case "d":
		fmt.Println(Divide(arg1, arg2))
	case "m":
		fmt.Println(Multiply(arg1, arg2))
	}
}
