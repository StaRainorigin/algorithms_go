package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	s := dummy
	p, q := list1, list2

	for p != nil && q != nil {
		if p.Val < q.Val {
			s.Next = p
			p = p.Next
		} else {
			s.Next = q
			q = q.Next
		}
		s = s.Next
	}

	if p != nil {
		s.Next = p
	} else {
		s.Next = q
	}

	return dummy.Next
}
