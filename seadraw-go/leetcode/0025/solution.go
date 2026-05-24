package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(pre *ListNode, k int) {
	cur := pre.Next
	for i := 0; i < k-1; i++ {
		nxt := cur.Next
		cur.Next = nxt.Next
		nxt.Next = pre.Next
		pre.Next = nxt
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for {
		p := pre
		for i := 0; i < k && p != nil; i++ {
			p = p.Next
		}
		if p == nil {
			break
		}

		reverse(pre, k)

		for range k {
			pre = pre.Next
		}
	}

	return dummy.Next
}
