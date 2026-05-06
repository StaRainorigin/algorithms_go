package solution

import (
	"slices"
)

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m := len(boxGrid)
	if m == 0 {
		return nil
	}
	n := len(boxGrid[0])

	ans := make([][]byte, n)
	for i := range n {
		ans[i] = make([]byte, m)
	}

	for i, row := range boxGrid {
		for j := range row {
			ans[j][i] = boxGrid[i][j]
		}
	}

	for i := range n {
		slices.Reverse(ans[i])
	}

	blocks := make([]int, m)
	for i := range m {
		blocks[i] = n - 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			switch ans[i][j] {
				case '*': {
					blocks[j] = i - 1
				}
				case '#': {
					ans[i][j] = '.'
					ans[blocks[j]][j] = '#'
					blocks[j]--
				}
			}
		}
	}

	return ans
}
