package solution

import (
	"math"
)

func finishInner(fstStartTime, fstDuration, secStartTime, secDuration []int) int {
	pre := math.MaxInt
	for i := range fstStartTime {
		pre = min(pre, fstStartTime[i]+fstDuration[i])
	}

	res := math.MaxInt
	for i := range secStartTime {
		res = min(res, max(secStartTime[i], pre)+secDuration[i])
	}

	return res
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	landFirst := finishInner(landStartTime, landDuration, waterStartTime, waterDuration)
	waterFirst := finishInner(waterStartTime, waterDuration, landStartTime, landDuration)
	return min(landFirst, waterFirst)
}
