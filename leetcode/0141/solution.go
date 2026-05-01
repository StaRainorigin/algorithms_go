package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p, q := head, head
	for q.Next != nil && q.Next.Next != nil {
		q = q.Next.Next
		p = p.Next
		if p == q {
			return true
		}
	}
	return false
}
