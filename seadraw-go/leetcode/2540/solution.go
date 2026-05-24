package solution

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	m, n := len(nums1), len(nums2)
	for i < m && j < n {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
