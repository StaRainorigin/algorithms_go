package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	var ans *TreeNode
// 	var hasNode func(*TreeNode) bool 
// 	hasNode = func(root *TreeNode) bool {
// 		if root == nil {
// 			return false
// 		}
// 		leftHas, rightHas := hasNode(root.Left), hasNode(root.Right)
// 		selfHas := root == p || root == q
// 		if (leftHas && rightHas) || (leftHas && selfHas) || (rightHas && selfHas) {
// 			ans = root
// 			return false
// 		}
// 		if selfHas || leftHas || rightHas {
// 			return true
// 		}
// 		return false
// 	}
// 	hasNode(root)
// 	return ans
// }
