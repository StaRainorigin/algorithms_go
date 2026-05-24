package solution

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	mp := map[*Node]*Node{}
	
	var dfs func(*Node) 
	dfs = func(node *Node) {
		_, ok := mp[node]
		if node == nil || ok {
			return 
		}
		mp[node] = &Node{node.Val, []*Node{}}
		for _, neighborNode := range node.Neighbors {
			dfs(neighborNode)
		}
	}
	dfs(node)

	for oldNode, newNode := range mp {
		for _, neighborNode := range oldNode.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, mp[neighborNode])
		}
	}

	return mp[node]
}
