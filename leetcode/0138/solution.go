package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for p := head; p != nil; p = p.Next.Next {
		q := &Node{p.Val, p.Next, nil}
		p.Next = q
	}

	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}

	dummy := &Node{}
	q := dummy
	for p := head; p != nil; p = p.Next{
		copyNode := p.Next
		p.Next = copyNode.Next
		q.Next = copyNode
		q = q.Next
	}

	return dummy.Next
}
