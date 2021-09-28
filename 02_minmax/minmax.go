package main

import (
	"fmt"
)

func main() {
	var a1, a2, min, max float64
	fmt.Println("This program calculates the maximum and minimum of two numbers.")
	fmt.Print("Enter the first number: ")
	fmt.Scan(&a1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&a2)

	if a1 < a2 {
		min = a1
		max = a2
	} else {
		min, max = a2, a1
	}

	fmt.Println("Maximum =", max)
	fmt.Println("Minimum =", min)
}
