package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	prev := math.MinInt
	
	var inorder func(*TreeNode) bool 
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		
		left := inorder(root.Left)
		cur := prev < root.Val
		prev = root.Val
		return left && cur && inorder(root.Right)
	}
	
	return inorder(root)
}
