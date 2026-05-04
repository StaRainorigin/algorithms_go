package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}

	k %= n
	if k == 0 {
		return head
}

	dummy := &ListNode{Next: head}
	fast := dummy
	for ;k > 0 ;k-- {
		fast = fast.Next
	}

	slow := dummy
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	
	dummy.Next = slow.Next
	slow.Next = nil
	fast.Next = head

	return dummy.Next
}
