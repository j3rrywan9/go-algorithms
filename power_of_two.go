package main

import (
	"fmt"
)

// LCOJ No. 231
func powerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}

func main() {
	res := powerOfTwo(64)
	fmt.Printf("%t\n", res)
	res = powerOfTwo(121)
	fmt.Printf("%t\n", res)
}
