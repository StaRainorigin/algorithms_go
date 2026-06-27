package solution

func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	ans := 0
	for x := range cnt {
		if x == 1 { // 1的n次方都是自己，需要单独处理
			continue
		}
		cur := 0
		for ; cnt[x] >= 2; x *= x {
			cur += 2
		}
		if cnt[x] >= 1 {
			cur++
		} else {
			cur--
		}
		ans = max(ans, cur)
	}

	if c := cnt[1]; c > 0 {
		ans = max(ans, c+c%2-1)
	}
	return ans
}
