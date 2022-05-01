/*
Exercise 13.3

Writer three different implementations of a function that finds the greatest number within an array. Write one function that is O(n ^ 2), one that is O(n log n), and one that is O(n).

Quadratic time: 
	nested loops: use a local variable to keep track of highest found, and compare each element to every other element in the array one element at a time

Loglinear time: sort the array in descending order and return the first element

Linear time: use a local variable to keep track of highest number found while iterating through each element. After going through the entire array, return the value of highest number found variable
*/
package main

import (
	"fmt"
	"sort"
)

func findGreatestNumQuadratic(nums []int) int {
	highestNumFound := nums[0]

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > highestNumFound {
				highestNumFound = nums[i]
			}

			if nums[j] > highestNumFound {
				highestNumFound = nums[j]
			}
		}
	}

	return highestNumFound
}

func findGreatestNumLogLinear(nums []int) int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Sort(sort.Reverse(sort.IntSlice(numsCopy)))

	return numsCopy[0]
}

func findGreatestNumLinear(nums []int) int {
	highestNumFound := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > highestNumFound {
			highestNumFound = nums[i]
		}
	}

	return highestNumFound
}

func main() {
  fmt.Println(findGreatestNumQuadratic([]int{90, 1, 4, 5, 89, 99, 56, 14}))   // 99
  fmt.Println(findGreatestNumLogLinear([]int{90, 1, 4, 5, 89, 99, 56, 14}))   // 99
  fmt.Println(findGreatestNumLinear([]int{90, 1, 4, 5, 89, 99, 56, 14}))      // 99

  fmt.Println(findGreatestNumQuadratic([]int{1, 4, 5, 89, 99, 56, 14, 101}))  // 101
  fmt.Println(findGreatestNumLogLinear([]int{1, 4, 5, 89, 99, 56, 14, 101}))  // 101
  fmt.Println(findGreatestNumLinear([]int{1, 4, 5, 89, 99, 56, 14, 101}))     // 101
}
