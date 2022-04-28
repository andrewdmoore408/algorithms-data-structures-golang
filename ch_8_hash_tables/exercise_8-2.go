/*
Exercise 8.2

Write a function which takes an array of strings and returns the first duplicate value it finds, e.g., ["a", "b", "c", "d", "c", "e", "f"] should return "c", since it's the first duplicate value in the array. Assume one pair of duplicates exists

Runtime should be O(n)

input: array of strings
output: string (first duplicate value found)

algo:
	initialize a variable stringsFound to an empty map of [string]boolean
	iterate through input array
		if stringsFound contains current string, return current string
		set current string as a key and its value to true in stringsFound
*/

package main

import "fmt"

func getFirstDuplicateString(words []string) string {
	stringsFound := make(map[string]bool)

	for _, word := range words {
		if stringsFound[word] == true {
			return word
		}

		stringsFound[word] = true
	}

	return ""
}

func main() {
  words := []string{"a", "b", "c", "d", "c", "e", "f"}
	fmt.Println(getFirstDuplicateString(words)) // "c"

	words = []string{"a", "b", "c", "d", "e", "f", "a"}
  fmt.Println(getFirstDuplicateString(words)) // "a"

	words = []string{"apple", "banana", "cherry", "cherry", "date", "elderberry", "fig"}
  fmt.Println(getFirstDuplicateString(words)) // "cherry"

	words = []string{"apple", "banana", "cherry", "date", "elderberry", "banana"}
  fmt.Println(getFirstDuplicateString(words)) // "banana"
}

