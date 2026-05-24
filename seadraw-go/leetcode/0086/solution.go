package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	sml, big := &ListNode{}, &ListNode{}
	p, q := sml, big
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			p.Next = cur
			p = p.Next
		} else {
			q.Next = cur
			q = q.Next
		}
	}
	p.Next = big.Next
	q.Next = nil
	return sml.Next
}
