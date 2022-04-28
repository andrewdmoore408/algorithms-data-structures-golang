/*
Exercise 9.4

Write a function which uses a stack to reverse a string; for example, "abcde" would become "edcba"
*/

package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []string
}

func (s *Stack) push(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) pop() (string, bool) {
	if len(s.data) == 0 {
		return "", false
	}

  popped := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]

	return popped, true
}

func (s Stack) read() string {
	if len(s.data) == 0 {
		return ""
	} else {
		return s.data[len(s.data) - 1]
	}
}

func reverseString(toReverse string) string {
	stack := Stack{strings.Split(toReverse, "")}

	reversed := ""

	for ok := true; ok == true; {
		letter, ok := stack.pop()

		if ok {
			reversed += letter
		} else {
			break
		}
	}

	return reversed
}

func main() {
	fmt.Println(reverseString("abcde")) 			// edcba
	fmt.Println(reverseString("zyx"))			// xyz
	fmt.Println(reverseString(""))		  		// ""
	fmt.Println(reverseString(".desrever si txet sihT"))	// This text is reversed.
}

