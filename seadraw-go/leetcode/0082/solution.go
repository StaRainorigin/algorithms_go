package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next.Val
		if cur == pre.Next.Next.Val {
			for pre.Next != nil && pre.Next.Val == cur {
				pre.Next = pre.Next.Next
			}
		} else {
			pre = pre.Next 
		}
	}

	return dummy.Next
}
