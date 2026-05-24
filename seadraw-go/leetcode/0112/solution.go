package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var pot func(*TreeNode, int) bool
	pot = func(root *TreeNode, cur int) bool {
		if root == nil {
			return false
		}

		cur += root.Val
		if root.Left == nil && root.Right == nil {
			return cur == targetSum
		}
		return pot(root.Left, cur) || pot(root.Right, cur)
	}

	return pot(root, 0)
}
