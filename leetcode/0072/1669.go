func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    dum := &ListNode{Next: list1}
    p := dum
    for i := 0; i < a; i++ {
        p = p.Next
    }
    left := p

    for i := 0; i < b-a+1; i++ {
        p = p.Next
    }
    right := p.Next

    left.Next = list2
    for list2.Next != nil {
        list2 = list2.Next
    }
    list2.Next = right
    return dum.Next
}