package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	ans := 0

	var pot func(*TreeNode, int)
	pot = func(root *TreeNode, cur int) {
		if root == nil {
			return 
		}
		cur = cur*10 + root.Val
		if root.Left == nil && root.Right == nil {
			ans += cur
		}
		pot(root.Left, cur)
		pot(root.Right, cur)
	}

	pot(root, 0)
	return ans
}
