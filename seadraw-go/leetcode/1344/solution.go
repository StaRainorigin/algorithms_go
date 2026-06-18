package solution

import "math"

func angleClock(hour int, minutes int) float64 {
	d := math.Abs(float64(hour%12)*30-float64(minutes)*5.5)
	return min(d, 360-d)
}