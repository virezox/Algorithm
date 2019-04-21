package main

import "fmt"

func main() {
	a := "avb"
	fmt.Println(a[0:0])
}

func longestCommonPrefix(strs []string) (prefix string) {
	shortest := -1
	for _, s := range strs {
		if len(s) < shortest || shortest == -1 {
			shortest = len(s)
		}
	}

	for i := 1; i < shortest+1; i++ {
		for _, s := range strs {
			if strs[0][0:i] != s[0:i] {
				prefix = strs[0][0 : i-1]
				return
			}
		}
		prefix = strs[0][0:i]
	}
	return
}
