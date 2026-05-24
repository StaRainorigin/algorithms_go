package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	
	ans := math.MaxInt
	prev := -100000
	
	var inorder func(*TreeNode) 
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		ans = min(ans, root.Val - prev)
		prev = root.Val
		inorder(root.Right)
	}

	inorder(root)
	return ans
}
