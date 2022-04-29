/*
Exercise 11.2

Use recursion to write a function that accepts an array of numbers and returns a new array containing just the even numbers.

notes: assuming integers for simplicity

input: array of nums
output: array of nums (new array including only the even numbers)

subproblem: the array minus one of its elements

algo:
	base case:
		empty array: return
	recursing case:
		initialize newArray to an empty array of ints
		if first element is even
			add first element to newArray

		return newArray combined with elements from recursing with remainder of array
*/

package main

import (
	"fmt"
)

func getEvenNums(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	newSlice := make([]int, 0)

	if nums[0] % 2 == 0 {
		newSlice = append(newSlice, nums[0])
	}

	return append(newSlice, getEvenNums(nums[1:])...)
}

func main() {
	fmt.Println(getEvenNums([]int{1, 2, 3, 4, 5}))				  // [2, 4]
	fmt.Println(getEvenNums([]int{1, 3, 5, 7, 9}))				  // []
	fmt.Println(getEvenNums([]int{1, 2, 3, 4, 6, 10, 14, 15}))		  // [2, 4, 6, 10, 14]
}

