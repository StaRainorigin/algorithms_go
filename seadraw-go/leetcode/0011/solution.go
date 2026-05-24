package solution

func maxArea(height []int) int {
	res := 0
	i, j := 0, len(height) - 1
	for i < j {
		if height[i] < height[j] {
			res = max(res, (j - i) * height[i])
			i++
		} else {
			res = max(res, (j - i) * height[j])
			j--
		}
	}
	return res
}
