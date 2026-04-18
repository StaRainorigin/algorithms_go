package solution

func isHappy(n int) bool {
	m := map[int]struct{}{}
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
		x := 0
		for ; n > 0; n /= 10 {
			x += (n % 10) * (n % 10)
		}
		n = x
	}
	return true
}
