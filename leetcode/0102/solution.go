package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{}
	q = append(q, root)
	ans := [][]int{}

	for len(q) > 0 {
		layer := []int{}
		p := []*TreeNode{}
		for _, cur := range q {
			layer = append(layer, cur.Val)
			if cur.Left != nil {
				p = append(p, cur.Left)
			}
			if cur.Right != nil {
				p = append(p, cur.Right)
			}
		}
		q = p
		ans = append(ans, layer)
	}

	return ans
}
