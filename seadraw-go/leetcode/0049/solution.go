package solution

import (
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	cnt := map[string][]string{}
	for _, s := range strs {
		cur := []byte(s)
		slices.Sort(cur)
		sorted := string(cur)
		cnt[sorted] = append(cnt[sorted], s)
	}
	return slices.Collect(maps.Values(cnt))
}
