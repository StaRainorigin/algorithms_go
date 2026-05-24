package solution

import (
	"slices"
)

func minimumEffort(tasks [][]int) (e int) {
	slices.SortFunc(tasks, func(a, b []int) int {
		return (a[1] - a[0]) - (b[1] - b[0]) // 按照 minimum - actual 从小到大排序
	})

	for _, t := range tasks {
		actual, minimum := t[0], t[1]
		// 完成 t 之后的能量为 e，那么完成 t 之前的能量为 e+actual，同时该能量必须至少为 minimum
		e = max(e+actual, minimum)
	}
	return
}

