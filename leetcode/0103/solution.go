package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{}
	q = append(q, root)
	ans := [][]int{}
	operate := 0

	for len(q) > 0 {
		layer := []int{}
		p := []*TreeNode{}
		n := len(q)
		for i, cur := range q {
			if operate == 0 {
				layer = append(layer, cur.Val)
			} else {
				layer = append(layer, q[n-i-1].Val)
			}
			
			if cur.Left != nil {
				p = append(p, cur.Left)
			}
			if cur.Right != nil {
				p = append(p, cur.Right)
			}
		}
		ans = append(ans, layer)
		q = p
		operate = (operate + 1) % 2 
	}
	return ans
}
