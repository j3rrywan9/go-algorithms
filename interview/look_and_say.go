package interview

/*
Seed: 1
Sequence: 11, 21, 1211, 111221

Seed: 33
Sequence: 23, 1213, 11121113, 31123113

Write a function to take a seed, output the sequence after specified iterations:

func lookAndSay(seed, iterations int) string

For example:

lookAndSay(1, 4) = 111221
*/

import (
	"strconv"
)

func lookAndSay(seed, iterations int) string {
	// Convert seed (integer) to string
	input := strconv.Itoa(seed)
	// Return the input directly if the iterations is less than 1
	if iterations < 1 {
		return input
	}
	ret := ""
	prev, current := string(input[0]), string(input[0])
	for i := 0; i < iterations; i++ {
		//fmt.Printf("input of iteration %d, %s\n", i + 1, input)
		counter := 0
		for j := 0; j < len(input); j++ {
			current = string(input[j])
			if current == prev {
				counter += 1
				//fmt.Printf("counter: %d\n", counter)
			} else {
				ret = ret + strconv.Itoa(counter) + prev
				counter = 1
				prev = current
				//fmt.Printf("counter: %d\n", counter)
			}
			//fmt.Printf("prev: %s\n", prev)
			//fmt.Printf("cur: %s\n", current)
		}
		ret = ret + strconv.Itoa(counter) + prev
		input = ret
		prev, current = string(input[0]), string(input[0])
		ret = ""
	}
	return input
}
