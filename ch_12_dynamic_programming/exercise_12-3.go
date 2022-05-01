/*
Exercise 12.3

Here is a solution to the "Unique Paths" problem from an exercise in the previous chapter. Use memoization to improve its efficiency.

func uniquePaths(rows, columns int) int {
	if rows == 1 || columns == 1 { return 1 }

	return uniquePaths(rows - 1, columns) + uniquePaths(rows, columns - 1)
}
*/
package main

import (
	"fmt"
	_ "sort"
	"strings"
	"strconv"
)

func uniquePaths(rows, columns int, memo(map[string]int)) int {
  if rows == 1 || columns == 1 { return 1 }

	rowAndCol := make([]string, 0)
	rowAndCol = append(rowAndCol, strconv.Itoa(rows), strconv.Itoa(columns))
	key := strings.Join(rowAndCol, " ")

  numPaths, found := memo[key]

	if found {
		return numPaths
	} else {
		memo[key] = uniquePaths(rows - 1, columns, memo) + uniquePaths(rows, columns - 1, memo)
	}

	return memo[key]
}

func main() {
  // fmt.Println(uniquePaths(3, 3)) // 6
  // fmt.Println(uniquePaths(4, 2)) // 4
  // fmt.Println(uniquePaths(5, 6)) // 126

  fmt.Println(uniquePaths(3, 3, make(map[string]int))) // 6
  fmt.Println(uniquePaths(4, 2, make(map[string]int))) // 4
  fmt.Println(uniquePaths(5, 6, make(map[string]int))) // 126
}
