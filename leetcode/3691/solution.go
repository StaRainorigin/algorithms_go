package solution

import (
	"container/heap"
	"math/bits"
)

type pair struct{ min, max int }

func op(a, b pair) pair {
	return pair{min(a.min, b.min), max(a.max, b.max)}
}

type ST [][16]pair // 16 = bits.Len(5e4)

func newST(a []int) ST {
	n := len(a)
	w := bits.Len(uint(n))
	st := make(ST, n)
	for i, x := range a {
		st[i][0] = pair{x, x}
	}
	for j := 1; j < w; j++ {
		for i := range n - 1<<j + 1 {
			st[i][j] = op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// [l,r) 左闭右开
func (st ST) query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	p := op(st[l][k], st[r-1<<k][k])
	return p.max - p.min
}

func maxTotalValue(nums []int, k int) (ans int64) {
	n := len(nums)
	st := newST(nums)
	h := make(hp, n)
	for i := range h {
		h[i] = tuple{st.query(i, n), i, n} // 子数组极差，左端点，右端点加一
	}
	// 由于 h 是递减的，无需堆化

	for ; k > 0 && h[0].d > 0; k-- {
		ans += int64(h[0].d)
		h[0].r--
		h[0].d = st.query(h[0].l, h[0].r)
		heap.Fix(&h, 0)
	}
	return
}

type tuple struct{ d, l, r int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d > h[j].d } // 最大堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
