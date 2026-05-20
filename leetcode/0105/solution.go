package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	n := len(preorder)
// 	if n == 0 {
// 		return nil
// 	}

// 	leftSize := slices.Index(inorder, preorder[0])
// 	left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
// 	right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])

// 	return &TreeNode{preorder[0], left, right}
// }

func buildTree(preorder []int, inorder []int) *TreeNode {
	mp := map[int]int{}
	for i, x := range inorder {
		mp[x] = i
	}

	n, m := len(preorder), len(inorder)
	
	var buildTreeInner func(int, int, int, int) *TreeNode
	buildTreeInner = func(pl, pr, il, ir int) *TreeNode {
		if pl >= pr {
			return nil
		}

		root, _ := mp[preorder[pl]]
		leftSize := root - il
		left := buildTreeInner(pl+1, pl + 1 + leftSize, il, il + leftSize)
		right := buildTreeInner(pl + 1 + leftSize, pr, il + 1 + leftSize, ir)

		return &TreeNode{preorder[pl], left, right}
	}

	return buildTreeInner(0, n, 0, m)
}
