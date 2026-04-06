package solution

func removeDuplicates(nums []int) int {
    n := len(nums)
    res := 1
    if n > 1 {
    	i, j := 0, 1
	    for j < n {
	   		if nums[i] != nums[j] {
				i++
				nums[i] = nums[j]
				res++
			} 
			j++
	    }
    }
    return res
}