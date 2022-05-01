/*
Exercise 12.1

Function given below accepts an array of numbers and returns the sum (as long as it doesn't go over 100, in which case any number that would lead to that result is ignored). It makes unnecessary recursive calls: fix the code to eliminate unneeded recursing.
*/

package main

import "fmt"

/* 
Original function below: makes 2 recursive calls on each iteration, leading to O(2 ^ n) runtime.
*/
func addUntil100(nums []int) int {
	if len(nums) == 0 { return 0 }

	if nums[0] + addUntil100(nums[1:]) > 100 {
		return addUntil100(nums[1:])
	} else {
		return nums[0] + addUntil100(nums[1:])
	}
}

/*
	Fixed version of this function below. By using some local variables and only making a recursive call once within the function body, the runtime is reduced from O(2 ^ n) to O(n).
*/
func addUntil100Fixed(nums []int) int {
	if len(nums) == 0 { return 0 }

	firstNum, sumOfRemaining := nums[0], addUntil100Fixed(nums[1:])
	currentSum := firstNum + sumOfRemaining

	if currentSum > 100 {
		return sumOfRemaining
	} else {
		return currentSum
	}
}

func main() {
  fmt.Println(addUntil100([]int{1, 2, 3, 4, 5}))            // 15
  fmt.Println(addUntil100([]int{10, 20, 30, 40, 50}))       // 100 (max return)
  fmt.Println(addUntil100Fixed([]int{1, 2, 3, 4, 5}))       // 15
  fmt.Println(addUntil100Fixed([]int{10, 20, 30, 40, 50}))  // 100 (max return)
  
  fmt.Println(addUntil100([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50}))       // 100 (max return)
  fmt.Println(addUntil100Fixed([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50}))  // 100 (max return)
}

