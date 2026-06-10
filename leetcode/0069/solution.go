package solution

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	// 初始猜测值可以是 x/2 或 x
	y := x
	for y > x/y {
		y = (y + x/y) / 2
	}

	return y
}
