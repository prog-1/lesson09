package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Type `int` represents integer numbers,
	// e.g. -2, -1, 0, 1, 2, ...
	// var n, m int
	// Assigning values to each variable individually:
	// m = 5
	// n = 7
	// Can also assign values for multiple variables at the same type.
	// m, n = 5, 7
	// A better alternative:
	m, n := 5, 7 // type for `m` and `n` gets derived automatically.

	// Conditions
	// Condition is an expression that evaluates/results in true or false.
	fmt.Println(m < n) // true
	// Comparison operators: <, ==, <=, !=, >, >=
	// '=' and '==' are different operators!

	// Integer variable can be compared to an integer literal (even when the
	// literal contains a decimal separator).
	fmt.Println(m < 3.0) // false
	// fmt.Println(n < 3.3) <- wrong
	x := 3.3
	fmt.Println(n < int(x))     // false
	fmt.Println(float64(n) < x) // false
	// fmt.Println(n < math.Round(3.3)) <- wrong
	fmt.Println(n < int(math.Round(6.3))) // false
	fmt.Println(5 < 5.3)                  // true

	// String comparison.
	// Strings are ordered in a lexicographical order:
	// https://en.wikipedia.org/wiki/Lexicographical_order
	fmt.Println("a" < "b")                   // true
	fmt.Println("az" < "b")                  // true
	fmt.Println("a" == "A")                  // false. Comparison is case-sensitive
	fmt.Println("a" == strings.ToLower("A")) // true. The right string is converted to lower-case. There is also strings.ToUpper
	fmt.Println("🐴" == "🦄")                  // false
	// Strings and integers cannot be compared.
	// fmt.Println("5" == 5) <- wrong.

	// Logical operators.
	// || -- or
	// && -- and
	// ! -- not
	// 0 <= n < 10
	fmt.Println(0 <= n && n < 10)
	fmt.Println(!(n < 0) || !(n >= 10))
}
