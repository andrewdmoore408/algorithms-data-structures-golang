/*
Exercise 11.1

Write a recursive function that takes an array of strings and returns the total number of chars across all the strings. For example: if input is ["ab", "c", "def", "ghij"], output should be 10 since there are 10 letters total.

input: array of strings
output: integer (number of characters across all strings)

notes: assuming alphanumeric characters for purposes of this exercise.

algo:
	base case: 
		array has length of 1: return length of string
	recursing case:
		return length of first element plus the recursion of rest of the array
*/

package main

import (
	"fmt"
)

func countAllChars(strs []string) int {
	if len(strs) == 1 {
		return len(strs[0])
	} else {
		return len(strs[0]) + countAllChars(strs[1:])
	}
}

func main() {
	fmt.Println(countAllChars([]string{"ab", "c", "def", "ghij"})) 		// 10
	fmt.Println(countAllChars([]string{"abcde", "xyz"}))	 		// 8
	fmt.Println(countAllChars([]string{"this is a long string", ""}))       // 21
	fmt.Println(countAllChars([]string{"a", "b", "c", "x", "y", "z"}))      // 6
}

