package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right - left; i++ {
		nxt := cur.Next
		cur.Next = nxt.Next
		nxt.Next = pre.Next
		pre.Next = nxt
	}

	return dummy.Next
}
