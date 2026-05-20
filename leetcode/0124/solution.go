package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lVal := dfs(node.Left)
		rVal := dfs(node.Right)
		ans = max(ans, lVal + rVal + node.Val)
		return max(max(lVal, rVal) + node.Val, 0)
	}
	dfs(root)
	return ans
}
