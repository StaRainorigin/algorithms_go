package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := head
	for range k {
		fast = fast.Next
	}
	
	slow := dummy
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	
	return dummy.Next
}
