package solution

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	n := len(words)
	i := 0

	for i < n {
		// 1) 贪心选本行 [i, j)
		j := i
		wordLenSum := 0
		for j < n && wordLenSum+len(words[j])+(j-i) <= maxWidth {
			wordLenSum += len(words[j])
			j++
		}

		wordCount := j - i
		spaceTotal := maxWidth - wordLenSum
		lastLine := (j == n)

		var b strings.Builder

		// 2) 最后一行 or 单词数=1：左对齐
		if lastLine || wordCount == 1 {
			for k := i; k < j; k++ {
				if k > i {
					b.WriteByte(' ')
				}
				b.WriteString(words[k])
			}
			for b.Len() < maxWidth {
				b.WriteByte(' ')
			}
		} else {
			// 3) 普通行：均匀分配空格
			gaps := wordCount - 1
			each := spaceTotal / gaps
			extra := spaceTotal % gaps // 左边 extra 个缝多一个

			for k := i; k < j; k++ {
				b.WriteString(words[k])
				if k < j-1 {
					cnt := each
					if (k - i) < extra {
						cnt++
					}
					for t := 0; t < cnt; t++ {
						b.WriteByte(' ')
					}
				}
			}
		}

		res = append(res, b.String())
		i = j
	}

	return res
}
