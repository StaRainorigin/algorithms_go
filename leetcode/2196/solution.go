package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	mp := map[int]*TreeNode{}
	hasParent := map[int]bool{}

	for _, dst := range descriptions {
		hasParent[dst[1]] = true

		_, ok := mp[dst[0]]
		if !ok {
			mp[dst[0]] = &TreeNode{Val: dst[0]}
		}
		_, ok = mp[dst[1]]
		if !ok {
			mp[dst[1]] = &TreeNode{Val: dst[1]}
		}

		if dst[2] == 0 {
			mp[dst[0]].Right = mp[dst[1]]
		} else {
			mp[dst[0]].Left = mp[dst[1]]
		}
	}

	for _, dst := range descriptions {
		if !hasParent[dst[0]] {
			return mp[dst[0]]
		}
	}

	return nil
}
