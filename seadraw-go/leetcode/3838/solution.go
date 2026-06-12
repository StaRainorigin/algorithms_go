package solution

import "strings"

func mapWordWeights(words []string, weights []int) string {
	var ans strings.Builder
	for _, word := range words {
		sum := 0
		for _, c := range word {
			sum += weights[c-'a']
		}
		sum %= 26
		ans.WriteRune('z'-rune(sum))
	}
	return ans.String()
}
