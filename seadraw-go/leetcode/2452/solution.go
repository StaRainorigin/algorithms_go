package solution

func twoEditWords(queries []string, dictionary []string) []string {
	ans := []string{}
	for _, q := range queries {
		for _, d := range dictionary {
			count := 0
			if len(q) != len(d) {
				continue
			}
			for i := range q {
				if q[i] != d[i] {
					count++
					if count > 2{
						break
					}
				}
			}
			if count <= 2{
				ans = append(ans, q)
				break
			}
		}
	}
	return ans
}
