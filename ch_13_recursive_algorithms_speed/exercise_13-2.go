/*
Exercise 13.2

The below function finds the "missing number" from an array of integers. that is, the array is expected to have all integers from 0 up to the array's length, but one is missing.

For example: the array [5, 2, 4, 1, 0] is missing the number 3, and the array [9, 3, 2, 5, 6, 7, 1, 0, 4] is missing the number 8.

This example function is O(n ^ 2) (the includes method alone is already O(n), since the computer needs to search the entire array to find n):

func includes(haystack []int, needle int) bool {
	for _, num := range haystack {
		if num == needle {
			return true
		}
	}

	return false
}

func findMissingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if !includes(nums, i) {
			return i
		}
	}

	// returning a negative number if no missing number is found
	return -1
}

Use sorting to write a new implementation of this function that only takes O(n log n). (There are even faster implementations, but we're focusing on using sorting as a technique to make code faster).
*/
package main

import (
	"fmt"
	"sort"
)

func findMissingNumber(nums []int) int {
  numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(numsCopy)

	for i := 0; i < len(numsCopy); i++ {
		if numsCopy[i] != i {
			return i
		}
	}

  // returning a negative number if no missing number is found
  return -1
}

func main() {
  fmt.Println(findMissingNumber([]int{2, 3, 1, 4, 5, 0, 7}))   // 6
	fmt.Println(findMissingNumber([]int{0, 1, 2, 3, 5, 6}))      // 4
	fmt.Println(findMissingNumber([]int{6, 5, 4, 1, 0, 2}))      // 3
	fmt.Println(findMissingNumber([]int{7, 0, 2, 4, 5, 3}))      // 1
}

