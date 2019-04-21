package main

import (
	"fmt"
	"math"
)

func main() {
	a := -32167261
	fmt.Println(reverse(a))
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)
}

func reverse(x int) int {
	r, a := 0, 0
	for x != 0 {
		r = x % 10
		r += a * 10
		a = r
		x /= 10
	}
	if r < math.MinInt32 || r > math.MaxInt32 {
		r = 0
	}

	return r
}
