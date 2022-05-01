/*
Exercise 13.1

Given an array of positive numbers, write a function that returns the greatest product of any three numbers. The approach of using three nested loops would clock in at O(n ^ 3), which is very slow. use sorting to implement the function in a way that it computes at O(n log n) speed. (There are even faster implementatations, but we're focusing on using sorting as a technique to make code faster).

input: array (of positive numbers)
output: integer (greatest product of any three nums from input)

note: going to assume integers for simplicity

algo:
	sort input array numerically
	return the product of the rightmost numbers in the sorted array
*/
package main

import (
	"fmt"
	"sort"
)

func greatestProductOf3(nums []int) int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(numsCopy)
 
	product := 1

	for index := len(numsCopy) - 3; index < len(numsCopy); index++ {
		product *= numsCopy[index]
	}

	return product
}

func main() {
        fmt.Println(greatestProductOf3([]int{1, 5, 10, 4, 3, 8, 3, 4, 2}))       // 400
	fmt.Println(greatestProductOf3([]int{10, 8, 5, 3, 4, 5, 1, 1, 2}))       // 400
	fmt.Println(greatestProductOf3([]int{10, 11, 12, 1, 2, 3, 4}))           // 1320
	fmt.Println(greatestProductOf3([]int{90, 101, 1, 2, 3, 4, 5, 120}))      // 1090800
}

