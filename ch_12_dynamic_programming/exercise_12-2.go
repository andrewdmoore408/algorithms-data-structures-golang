/*
Exercise 12.2

The following function uses recursion to calculate the nth number from a mathematical sequence known as the "Golomb" sequence. It's terribly inefficient though! Use memoization to optimize it.
*/

package main

import "fmt"

/* 
Original function below: makes 3 recursive calls on each invocation (except for the base case), leading to O(3 ^ n) runtime.
*/
func golomb(n int) int {
	if n == 1 { return 1 }

	return 1 + golomb(n - golomb(golomb(n - 1)))
}

/*
	Fixed version of this function below. By using memoization, each value needs to be calculated only one time, leading to a dramatic improvement in runtime.
*/

func golombFixed(n int, memo map[int]int) int {
	if n == 1 { return 1 }

	if golombValue, wasFound := memo[n]; wasFound {
		return golombValue
	} else {
		memo[n] = 1 + golombFixed(n - golombFixed(golombFixed(n - 1, memo), memo), memo);
	}

	return memo[n]
}

func main() {
  fmt.Println(golomb(49))
  fmt.Println(golombFixed(49, make(map[int]int)))
}

