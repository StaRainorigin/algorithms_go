package solution

import "strings"

func simplifyPath(path string) string {
	st := []string{}
	for s := range strings.SplitSeq(path, "/") {
		if s == "" || s == "." {
            continue
        }
        if s != ".." {
            st = append(st, s)
        } else if len(st) > 0 {
            st = st[:len(st)-1]
        }
	}
	return "/" + strings.Join(st, "/")
}
