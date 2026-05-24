package solution

import "strconv"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	has := map[int]struct{}{}
	for _, x := range arr1 {
		for ; x > 0; x /= 10 {
			has[x] = struct{}{}
		}
	}

	ans := 0
	for _, x := range arr2 {
		for ; x > 0; x /= 10 {
			_, ok := has[x]
			if ok {
				break
			}
		}
		ans = max(ans, x)
	}

	if ans == 0 {
		return 0
	}

	return len(strconv.Itoa(ans))
}
