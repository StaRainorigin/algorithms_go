package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var inorder func(*TreeNode) int
	inorder = func(root *TreeNode) int {
		if root == nil {
			return -1
		}

		if left := inorder(root.Left); left != -1 {
			return left
		}
		if k--; k == 0 {
			return root.Val
		}
		return inorder(root.Right)
	}

	return inorder(root)
}
