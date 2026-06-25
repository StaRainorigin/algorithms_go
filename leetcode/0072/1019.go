func nextLargerNodes(head *ListNode) (ans []int) {
    st := []int{}
    var f func(*ListNode, int)
    f = func(node *ListNode, i int) {
        if node == nil {
            ans = make([]int, i) 
            return
        }
        f(node.Next,i+1)
        l:=len(st)
        for l != 0 && st[l-1] <=node.Val{
            st = st[:l-1]
            l--
        }
        if l>0{
            ans[i] = st[l-1]
        }
        st = append(st,node.Val)

    }

    f(head, 0)
    return

}