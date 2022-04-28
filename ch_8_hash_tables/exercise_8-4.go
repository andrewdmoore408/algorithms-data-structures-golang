/*
Exercise 8.4

Write a function which returns the first non-duplicated character in a string: for example, "minimum" as input should return "n", since it's the first character found in the string which only occurs once.

Runtime should be O(n)

input: string
output: string

algo:
	initialize a variable lettersFound to a map of [string]int
	initialize a variable to the slice of individual letters from input
	go through each char in input and increment the corresponding value in lettersFound (using current char as key)
	iterate again through each char in input
		if lettersFound value for currentChar key is 1, return current char
	return an empty string
*/

package main

import (
	"fmt"
	"strings"
)

func getFirstUniqueChar(str string) string {
	lettersFound := make(map[string]int)
	splitLetters := strings.Split(str, "")

	for _, letter := range splitLetters {
		lettersFound[letter]++
	}

	for _, letter := range splitLetters {
		if lettersFound[letter] == 1 {
			return letter
		}
	}

	return ""
}	

func main() {
  str := "minimum"
	fmt.Println(getFirstUniqueChar(str)) // "n"

  str = "battery"
  fmt.Println(getFirstUniqueChar(str)) // "b"

	str = "kayak"
  fmt.Println(getFirstUniqueChar(str)) // "y"

	str = "structs"
  fmt.Println(getFirstUniqueChar(str)) // "r"
}

