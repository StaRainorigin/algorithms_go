func removeNodes(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    nextValid := removeNodes(head.Next)
    
    if head.Val < nextValid.Val {
        return nextValid
    }
    
    head.Next = nextValid
    return head
}