package solution

func reverse(x int) int {
	rev := 0
	for ; x > 0; x /= 10 {
		rev = rev * 10 + x % 10
	}
	return rev
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func mirrorDistance(n int) int {
	return abs(n - reverse(n))
}
