package solution

func sumAndMultiply(n int) int64 {
	sum, x := 0, 0
	for ; n > 0; n /= 10 {
		y := n % 10
		sum += y
		if y > 0 {
			x = x * 10 + y
		}
	}

	for ; x > 0; x /= 10 {
		n = n * 10 + x % 10
	}

	return int64(sum) * int64(n)
}
