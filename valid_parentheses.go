package main

import (
	"fmt"
)

func stringInSlice(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func pop(s *[]string) string {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func isValid(s string) bool {
	var stack []string
	var dict = map[string]string{"]": "[", "}": "{", ")": "("}
	dictKeys := make([]string, len(dict))
	dictValues := make([]string, len(dict))
	i := 0
	for k, v := range dict {
		dictKeys[i] = k
		dictValues[i] = v
		i++
	}
	for _, char := range s {
		if stringInSlice(string(char), dictValues) {
			stack = append(stack, string(char))
		} else if stringInSlice(string(char), dictKeys) {
			if stack == nil || dict[string(char)] != pop(&stack) {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s := "(){}"
	fmt.Printf("%t\n", isValid(s))
}
