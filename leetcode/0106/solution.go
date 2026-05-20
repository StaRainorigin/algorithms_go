package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func buildTree(inorder []int, postorder []int) *TreeNode {
// 	n := len(postorder)
// 	if n == 0 {
// 		return nil
// 	}

// 	leftSize := slices.Index(inorder, postorder[n-1])
// 	left := buildTree(inorder[:leftSize], postorder[:leftSize])
// 	right := buildTree(inorder[leftSize+1:], postorder[leftSize:n-1])

// 	return &TreeNode{postorder[n-1], left, right}
// }

func buildTree(inorder []int, postorder []int) *TreeNode {
	cnt := map[int]int{}
	for i, x := range inorder {
		cnt[x] = i
	}
	m, n := len(inorder), len(postorder)
	
	
	var buildTreeInner func(int, int, int, int) *TreeNode
	buildTreeInner = func(il, ir, pl, pr int) *TreeNode {
		if pl >= pr {
			return nil
		}

		rootIndex := cnt[postorder[pr-1]]
		leftSize := rootIndex - il
		left := buildTreeInner(il, il + leftSize, pl, pl + leftSize)
		right := buildTreeInner(il + leftSize + 1, ir, pl + leftSize, pr - 1)

		return &TreeNode{postorder[pr-1], left, right}
	}

	return buildTreeInner(0, m, 0, n)
}
