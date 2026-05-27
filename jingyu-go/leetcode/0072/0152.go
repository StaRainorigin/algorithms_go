func maxProduct(nums []int) int {
    ma,mi:= make([]int,len(nums)),make([]int,len(nums))
    ma[0],mi[0] = nums[0],nums[0]
    ans := nums[0]
    for i:= 1;i<len(nums);i++{
        ma[i] = max(ma[i-1]*nums[i],nums[i],mi[i-1]*nums[i])
        mi[i] = min(mi[i-1]*nums[i],nums[i],ma[i-1]*nums[i])
        ans = max(ans,ma[i])
    }
    return ans
}