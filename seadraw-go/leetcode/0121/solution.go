package solution

func maxProfit(prices []int) int {
	n := len(prices)
	res := 0
	if n < 2 {
		return res
	}
	curMax := prices[n-1]

	for i := n - 2; i >= 0; i-- {
		if prices[i] > curMax {
			curMax = prices[i]
		} else {
			res = max(res, curMax - prices[i])
		}
	}
	
	return res
}
