package solution

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	carry := 0
	head := ListNode{0, nil}
	s := &head
	for p != nil || q != nil || carry == 1 {
		pv, qv := 0, 0
		if p != nil {
			pv = p.Val
			p = p.Next
		}
		if q != nil {
			qv = q.Val
			q = q.Next
		}
		cur := pv + qv + carry
		s.Next = &ListNode{cur%10, nil}
		s = s.Next
		carry = cur/10
	}
	return head.Next
}
