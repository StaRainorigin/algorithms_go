package solution

func sequentialDigits(low, high int) (ans []int) {
	x0 := 12 // 第一个窗口
	pow10 := 10
	for length := 2; x0 <= high; length++ {
		pow10 *= 10
		x := x0
		for i := length; i <= 9 && x <= high; i++ {
			if x >= low {
				ans = append(ans, x)
			}
			// 窗口向右滑动，i+1 进入窗口，i+1-length 离开窗口
			x = x*10 + i + 1 - (i+1-length)*pow10
		}
		x0 = x0*10 + length + 1
	}
	return
}