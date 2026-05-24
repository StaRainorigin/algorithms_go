package solution

import "strings"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func furthestDistanceFromOrigin(moves string) int {
	l := strings.Count(moves, "L")
	r := strings.Count(moves, "R")
	return abs(r - l) + len(moves) - l - r
}
