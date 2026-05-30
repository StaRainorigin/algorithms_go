package solution

import (
	"slices"
)

type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

type uf []int

func (f uf) find(x int) int {
	if f[x] != x {
		f[x] = f.find(f[x])
	}
	return f[x]
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	pos := []int{0}
	for _, q := range queries {
		m = max(m, q[1])
		if q[0] == 1 {
			pos = append(pos, q[1])
		}
	}
	m++

	left := make(uf, m+1)
	right := make(uf, m+1)
	for i := range left {
		left[i] = i
		right[i] = i
	}
	t := make(fenwick, m)
	slices.Sort(pos)
	for i := 1; i < len(pos); i++ {
		p, q := pos[i-1], pos[i]
		t.update(q, q-p)
		for j := p + 1; j < q; j++ {
			left[j] = p // 删除 j
			right[j] = q
		}
	}
	for j := pos[len(pos)-1] + 1; j < m; j++ {
		left[j] = pos[len(pos)-1] // 删除 j
		right[j] = m
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		x := q[1]
		pre := left.find(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			left[x] = x - 1 // 删除 x
			right[x] = x + 1
			nxt := right.find(x)   // x 右侧最近障碍物的位置
			t.update(nxt, nxt-pre) // 更新 d[nxt] = nxt - pre
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.preMax(pre), x-pre)
			ans = append(ans, maxGap >= q[2])
		}
	}
	slices.Reverse(ans)
	return
}