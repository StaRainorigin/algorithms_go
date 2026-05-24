package solution

import (
	"slices"
	"math"
	"math/bits"
	"sort"
)

type seg []int

func (t seg) update(node, l, r, i, val int) {
	if l == r { // 叶子
		t[node] = val
		return
	}
	m := (l + r) / 2
	if i <= m { // i 在左子树
		t.update(node*2, l, m, i, val)
	} else { // i 在右子树
		t.update(node*2+1, m+1, r, i, val)
	}
	t[node] = max(t[node*2], t[node*2+1])
}

func (t seg) query(node, l, r, ql, qr int) int {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return t[node]
	}
	m := (l + r) / 2
	if qr <= m { // [ql, qr] 在左子树
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 在右子树
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	return max(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func maximumJumps(nums []int, target int) int {
	// 排序去重，便于离散化
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	n := len(nums)
	m := len(sorted)

	t := make(seg, 2<<bits.Len(uint(m-1))) // 值域线段树
	for i := range t {
		t[i] = math.MinInt
	}

	// nums[0] 对应的 f[0] = 0
	t.update(1, 0, m-1, sort.SearchInts(sorted, nums[0]), 0)

	for j := 1; ; j++ {
		l := sort.SearchInts(sorted, nums[j]-target)       // >= nums[j]-target 的第一个数
		r := sort.SearchInts(sorted, nums[j]+target+1) - 1 // <= nums[j]+target 的最后一个数
		// t.query 返回满足 nums[j]-target <= nums[i] <= nums[j]+target 的最大的 f[i]
		fj := t.query(1, 0, m-1, l, r) + 1
		if j == n-1 {
			if fj < 0 {
				return -1
			}
			return fj
		}
		t.update(1, 0, m-1, sort.SearchInts(sorted, nums[j]), fj)
	}
}
