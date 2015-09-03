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
	r := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return r
}

func isValid(s string) bool {
	var stack []string
	var dict = map[string]string{"]":"[", "}":"{", ")":"("}
	dict_keys := make([]string, len(dict))
	dict_values := make([]string, len(dict))
	i := 0
	for k, v := range dict {
		dict_keys[i] = k
		dict_values[i] = v
		i++
	}
	for _, char := range s {
		if stringInSlice(string(char), dict_values) {
			stack = append(stack, string(char))
		} else if stringInSlice(string(char), dict_keys) {
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
	mystr := "(){}"
	fmt.Printf("%t\n", isValid(mystr))	
}

