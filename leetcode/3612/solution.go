package solution

import (
	"slices"
)

func processStr(s string) string {
	ans := []byte{}
	for _, x := range s {
		switch x {
			case '*': {
				if len(ans) > 0 {
					ans = ans[:len(ans)-1]
				}
			}
			case '#': {
				ans = append(ans, ans...)
			}
			case '%': {
				slices.Reverse(ans)
			}
			default: {
				ans = append(ans, byte(x))
			}
		}
	}
	return string(ans)
}