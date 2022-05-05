/*
Exercise 10.4

Write a recursive function which takes an arbitrarily-nested array (see below) which contains both numbers and arrays (which themselves can contain numbers and/or arrays).

The function should print out each number (and only numbers, not arrays).
*/

package main

import (
	"fmt"
)

type nested []interface{}

func printNestedValues(n nested) {
	for _, item := range n {
		switch element := item.(type) {
		case int:
			fmt.Println(element)
		case nested:
			printNestedValues(element)
		}
	}
}

func main() {
	is := make(nested, 0)
	is = append(is, 1, 2, 3)
	is = append(is, nested{4, 5, 6})
	is = append(is, 7)
	is = append(is, nested{8, nested{9, 10, 11, nested{12, 13, 14}}})

	fmt.Println(is)
	printNestedValues(is)
}
