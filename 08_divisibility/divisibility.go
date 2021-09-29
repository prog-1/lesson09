package main

import "fmt"

func main() {
	var dividend, divisor int
	fmt.Println("This program checks if one integer is divisible by another integer.")
	fmt.Print("Enter a dividend: ")
	fmt.Scan(&dividend)
	fmt.Print("Enter a divisor: ")
	fmt.Scan(&divisor)

	if divisor == 0 {
		fmt.Println(" Division by zero is an invalid operation :( ")
	} else {
		quotient := dividend / divisor
		remainder := dividend % divisor
		if remainder == 0 {
			fmt.Println(dividend, "is divisible by", divisor)
			fmt.Println(dividend, ":", divisor, "=", quotient)
		} else { // is not divisible (the remainder is not equal to 0)
			fmt.Println(dividend, "is not divisible by", divisor)
			fmt.Println(dividend, ":", divisor, "=", quotient, "and", remainder, "is a remainder")
		}
	}
}
