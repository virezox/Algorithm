package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(isPalindrome(a))
}

// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}

// 	v := []int{}
// 	for x > 0 {
// 		i := x % 10
// 		v = append(v, i)
// 		x /= 10
// 	}
// 	for a, b := 0, len(v)-1; a < len(v) && b >= 0; a, b = a+1, b-1 {
// 		if v[a] != v[b] {
// 			return false
// 		}
// 	}
// 	return true
// }

func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}

// func isPalindrome(x int) bool {
// 	s := strconv.Itoa(x)
// 	for i := 0; i < len(s); i++ {
// 		if s[i] != s[len(s)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }
