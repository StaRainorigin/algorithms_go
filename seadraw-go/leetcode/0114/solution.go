package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	dummy := &TreeNode{}
	cur := dummy

	var pot func(*TreeNode)
	pot = func(node *TreeNode) {
		if node == nil {
			return
		}
		left, right := node.Left, node.Right

		cur.Right = node
		cur = cur.Right
		node.Left = nil

		pot(left)
		pot(right)
	}

	pot(root)
}

