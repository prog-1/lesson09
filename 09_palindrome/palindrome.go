package main

import "fmt"

func main() {
	var n int
	// Integers < 100 are not accepted, because we cannot determine a preceeding
	// 0, e.g.
	// 10 and 010
	// 0, 00 and 000.
	fmt.Println("The program checks if a number from range [100, 999] is a palindrome.")
	fmt.Print("Enter a number from range [100, 999]: ")
	fmt.Scan(&n)
	if n < 100 || n > 999 {
		fmt.Println("The number must be in range [100, 999].")
		return
	}
	firstDigit := n / 100
	lastDigit := n % 10
	if firstDigit == lastDigit {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
