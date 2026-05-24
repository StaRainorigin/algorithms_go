package solution

import "strconv"

func evalRPN(tokens []string) int {
	st := []int{}
	for _, token := range tokens {
		x, err := strconv.Atoi(token)
		if err == nil {
			st = append(st, x)
			continue
		}
		num := st[len(st)-1]
		st = st[:len(st)-1]
		n := len(st)
		switch token {
			case "+": {
				st[n-1] += num
			}
			case "-": {
				st[n-1] -= num
			}
			case "*": {
				st[n-1] *= num
			}
			case "/": {
				st[n-1] /= num
			}
		}
	}
	return st[0]
}
