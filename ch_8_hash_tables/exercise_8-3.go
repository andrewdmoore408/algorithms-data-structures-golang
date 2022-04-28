/*
Exercise 8.3

Write a function which takes a string with all letters in the alphabet minus one--the function should return the missing letter. For example, "the quick brown box jumps over a lazy dog" contains all letters except "f", so it returns "f".

Runtime should be O(n)

input: string (containing 25 letters)
output: string (the one letter which wasn't present in the input)

algo:
	initialize a constant LETTERS to all letters in the latin alphabet
	initialize a variable lettersFound to a map of [string]bool
	go through each letter in input
		set lowercase version of that letter as a key in lettersFound and its value as true
	go through each letter in LETTERS
		if current letter is not in lettersFound, return current letter
*/

package main

import (
	"fmt"
	"strings"
)

func getMissingLetter(str string) string {
	const LETTERS = "abcdefghijklmnopqrstuvwxyz"
	lettersFound := make(map[string]bool)

	for _, letter := range strings.Split(str, "") {
		lettersFound[letter] = true
	}

	for _, letter := range strings.Split(LETTERS, "") {
		if !lettersFound[letter] {
			return letter
		}
	}

	return ""
}	

func main() {
  missingOneLetter := "the quick brown box jumps over a lazy dog"
	fmt.Println(getMissingLetter(missingOneLetter)) // "f"

	missingOneLetter = "bcdefghijklmnopqrstuvwxyz"
  fmt.Println(getMissingLetter(missingOneLetter)) // "a"

	missingOneLetter = "abcdefghijklmnopqrstuvwxy"
  fmt.Println(getMissingLetter(missingOneLetter)) // "z"

	missingOneLetter = "abcdefghijklnopqrstuvwxyz"
  fmt.Println(getMissingLetter(missingOneLetter)) // "m"
}

