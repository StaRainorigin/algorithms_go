package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	ans := []float64{}
	q := []*TreeNode{}
	q = append(q, root)
	
	for len(q) > 0 {
		avg := float64(0)
		nFloat := float64(len(q))
		p := []*TreeNode{}
		for _, cur := range q {
			avg += float64(cur.Val)
			if cur.Left != nil {
				p = append(p, cur.Left)
			}
			if cur.Right != nil {
				p = append(p, cur.Right)
			}
		}
		q = p
		ans = append(ans, avg / nFloat)
	}

	return ans
}
