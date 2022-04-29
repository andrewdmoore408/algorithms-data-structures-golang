/*
Exercise 11.4

Use recursion to write a function that accepts a string and returns the first index that contains the character "x". For example, the string "abcdefghijklmnopqrstuvwxyz" has an "x" at index 23. To keep things simple, assume the string definitely has at least one "x".

input: string
output: integer (the index at which "x" first appears)

notes: assume at least one "x" is present

algo:
	base case:
		first char in string is an "x"
		return 0
	recursing case:
		return 1 + value of recursing on the rest of the string (starting at index 1)
*/

package main

import (
	"fmt"
	_ "strings"
)

func getIndexFirstX(str string) int {
	if string(str[0]) == "x" {
		return 0
	}

	return 1 + getIndexFirstX(str[1:])
}

func main() {
	fmt.Println(getIndexFirstX("abcdefghijklmnopqrstuvwxyz"))   // 23
	fmt.Println(getIndexFirstX("xyz"))                          // 0
	fmt.Println(getIndexFirstX("One animal is the fox"))        // 20
}

