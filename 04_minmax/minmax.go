package main

import (
	"fmt"
)

func main() {
	var min, max float64
	fmt.Println("This program calculates the maximum and minimum of two numbers.")
	fmt.Print("Enter the first number: ")
	fmt.Scan(&min)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&max)

	if max < min {
		min, max = max, min
	} else if max == min {
		fmt.Println("The numbers are equal.")
	}

	fmt.Println("Maximum =", max)
	fmt.Println("Minimum =", min)
}
