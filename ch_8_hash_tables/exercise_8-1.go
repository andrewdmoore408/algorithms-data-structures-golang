/*
Exercise 8.1

Write a function that returns intersection of two arrays (a third array which contains all values which occur in both arrays)

Runtime should be O(n)

input: two arrays
output: new array (containing all values which are found in both arrays)

algo:
	return empty array if either input has length 0
	go through first array and build a map containing true for each value found
	initialize intersection to empty array
	iterate through second array
		if current value is found in the map, add currentVal to intersection
	return intersection
*/

package main

import "fmt"

func intersection(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 { return []int{} }
	aValues := make(map[int]bool)

	for _, aVal := range a {
		aValues[aVal] = true
	}

	intersection := make([]int, 0)

	for _, bVal := range b {
		if aValues[bVal] {
			intersection = append(intersection, bVal)
		}
	}

	return intersection
}

func main() {
  first, second := []int{1, 2, 3, 4, 5}, []int{5, 2}
	fmt.Println(intersection(first, second)) // [5, 2]

	first, second = []int{0, 1, 2, 3, 4, 5, 6}, []int{7, 8, 9}
	fmt.Println(intersection(first, second)) // []

	first, second = []int{0, 1, 2, 3, 4, 5, 6}, []int{7, 8, 9, 0, 3}
  fmt.Println(intersection(first, second)) // [0, 3]

	first, second = []int{0, 1, 2, 3, 4, 5, 6}, []int{}
  fmt.Println(intersection(first, second)) // []
}

