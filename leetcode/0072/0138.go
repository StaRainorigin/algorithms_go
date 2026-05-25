func copyRandomList(head *Node) *Node {
    for head.Next.Next!= nil{
        n := head.Next
        head.Next = &Node{
            Val : head.Val,
            Next : n,
            Random : head.Random,
        }
    }
}
