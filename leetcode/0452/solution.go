package solution

import (
	"math"
	"slices"
)

func findMinArrowShots(points [][]int) (ans int) {
    slices.SortFunc(points, func(a, b []int) int { return a[1] - b[1] }) 
    pre := math.MinInt
    for _, p := range points {
        if p[0] > pre { 
            ans++
            pre = p[1] 
        }
    }
    return
}