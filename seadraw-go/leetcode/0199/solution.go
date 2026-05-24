package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	ans := []int{}
	
	q := []*TreeNode{}
	q = append(q, root)

	for len(q) > 0 {
		n := len(q)
		ans = append(ans, q[n-1].Val)
		p := []*TreeNode{}
		for _, cur := range q {
			if cur.Left != nil {
				p = append(p, cur.Left)
			}
			if cur.Right != nil {
				p = append(p, cur.Right)
			}
		}
		q = p
	}
	return ans
}
