/*
Exercise 11.5

"Unique Paths" problem: Let's say you have a grid of rows and columns. Write a function that accepts a number of rows and a number of columns, and calculates the number of possible "shortest" paths from the upper-leftmost square to the lower-rightmost square.

"Shortest" path means that at every step you either move one square to the right or one square downward.

Return the number of shortest paths.

input: 2 integers (number of rows and number of columns)
output: integer (number of shortest paths that exist from top left square to bottom right)

notes: shortest path means the lowest number of squares to move to get from start to finish

algo:
	base case:
		row or column is 1 (this means there's only one path)
			return 1
	recursing case:
		return value of recursing both on one fewer row and one fewer column
*/

package main

import (
	"fmt"
)

func numShortestPaths(rows, columns int) int {
	if rows == 1 || columns == 1 {return 1}

	return numShortestPaths(rows - 1, columns) + numShortestPaths(rows, columns - 1)
}

func main() {
	fmt.Println(numShortestPaths(2, 2))  // 2
	fmt.Println(numShortestPaths(3, 3))  // 6
	fmt.Println(numShortestPaths(2, 4))  // 4
	fmt.Println(numShortestPaths(4, 3))  // 10
}

