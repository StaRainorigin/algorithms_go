package solution

func maxDistance(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	i, j := m - 1, n - 1
	ans := 0
	for i >= 0 && j >= 0 {
		for i >= 0 && nums1[i] <= nums2[j] {
			ans = max(ans, j - i)
			i--
		}
		j--
		if i > j {
			i--
		}
	}
	return ans
}
