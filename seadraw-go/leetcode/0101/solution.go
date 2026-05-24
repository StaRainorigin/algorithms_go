package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSame(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	return (p.Val == q.Val) && isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSame(root.Left, root.Right)
}
