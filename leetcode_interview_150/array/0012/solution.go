package solution

import "strings"

var maps = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	var builder strings.Builder // 使用 strings.Builder
	i := 0
	for num > 0 {
		cur_key := keys[i]
		// 如果当前数字大于或等于当前罗马数字值，则进行处理
		if num >= cur_key { // 将 num / cur_key > 0 改为 num >= cur_key，逻辑更清晰
			for j := 0; j < num/cur_key; j++ {
				builder.WriteString(maps[cur_key]) // 使用 WriteString
			}
			num %= cur_key
		} else {
			// 如果当前罗马数字值太大，则尝试下一个较小的值
			i++
		}
	}
	return builder.String() // 返回构建器中的字符串
}
