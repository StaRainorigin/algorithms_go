package solution

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func canReach(arr []int, start int) bool {
	n := len(arr)
	var canReachInner func (int) bool
	canReachInner = func(i int) bool {
		if (0 > i || i >= n) || arr[i] < 0{
			return false
		}
		cur := arr[i]
 		if cur == 0 {
			return true
		}
		arr[i] = -arr[i]
		return canReachInner(i + cur) || canReachInner(i - cur)
	}
	return canReachInner(start)
}
