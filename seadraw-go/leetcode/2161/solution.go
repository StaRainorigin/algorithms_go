package solution

import "cmp"

func pivotArray(nums []int, pivot int) []int {
	arr := [3][]int{}
	for _, x := range nums {
		cmp := cmp.Compare(x, pivot) //cmp = -1, 0, 1; [cmp+1] = [0, 1, 2]
		arr[cmp+1] = append(arr[cmp+1], x)
	}
	return append(append(arr[0], arr[1]...), arr[2]...)
}
