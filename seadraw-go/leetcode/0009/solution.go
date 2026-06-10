package solution

func isPalindrome(x int) bool {
	y := 0
	for x := x; x > 0; x /= 10 {
		y = y * 10 + x % 10
	}
	return x == y
}
