package main

func main() {

}

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	prev := &ListNode{}
	prev.Next = head
	res := prev
	for prev.Next != nil && prev.Next.Next != nil {
		odd := prev.Next
		even := odd.Next
		prev.Next, even.Next, odd.Next = even, odd, even.Next
		prev = odd
	}
	return res.Next
}
