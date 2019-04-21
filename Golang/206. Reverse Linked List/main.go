package main

import "fmt"

func main() {
	a := 1
	b := 10
	c := 1000

	a, b, c = a+b, b+c, a+c
	fmt.Printf("value a: %d   value b: %d   value c :%d", a, b, c)
}

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
