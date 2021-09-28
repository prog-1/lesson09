package main

import (
	"fmt"
	"math"
)

func main() {
	var a1, a2, min, max float64
	fmt.Println("This program calculates the maximum and minimum of two numbers.")
	fmt.Print("Enter the first number: ")
	fmt.Scan(&a1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&a2)

	max = (a1 + a2 + math.Abs(a1-a2)) / 2
	min = a1 + a2 - max

	fmt.Println("Maximum =", max)
	fmt.Println("Minimum =", min)
}
