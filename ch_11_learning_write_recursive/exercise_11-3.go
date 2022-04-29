/*
Exercise 11.3

There's a numerical sequence known as "Triangular Numbers." The pattern begins as 1, 3, 6, 10, 15, 21 and continues onward with the nth number in the pattern, which is n plus the previous number, e.g., the 7th number is 28, which is 7 (n) plus the previous number in the sequence (21, which is the sequence's value for 6).

Write a recursive function that accepts an integer and returns the correct triangular series number for that input

input: int
output: int

algo:
	base case: 
		input is 1
			return 1
	recursing case:
		return current number plus recursing on current number - 1
*/

package main

import (
	"fmt"
)

func getTriangularNum(n int) int {
	if n == 1 {
		return 1
	}

	return n + getTriangularNum(n - 1)
}

func main() {
	fmt.Println(getTriangularNum(7))     // 28
	fmt.Println(getTriangularNum(3))     // 6
	fmt.Println(getTriangularNum(6))     // 21
}

