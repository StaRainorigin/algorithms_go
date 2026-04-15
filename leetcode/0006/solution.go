package solution

import "strings"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	n := len(s)
	var res strings.Builder
	step := 2 * (numRows - 1)
	for i := range numRows {
		for j := i; j < n; j += 2 * (numRows - 1) {
			res.WriteByte(s[j])
			if i != 0 && i != numRows - 1 && step > 0 && step + j < n {
				res.WriteByte(s[step + j])
			}
		}
		step -= 2
	}
	return res.String()
}

