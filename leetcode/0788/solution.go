package solution

func rotatedDigits(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if isGood(i) {
			ans++
		}
	}
	return ans
}

func isGood(x int) bool {
	diff := false
	for x > 0 {
		d := x % 10
		if d == 3 || d == 4 || d == 7 {
			return false
		}
		if d == 2 || d == 5 || d == 6 || d == 9 {
			diff = true
		}
		x /= 10
	}
	return diff
}
