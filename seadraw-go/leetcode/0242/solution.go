package solution

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    cnt := [26]int{}
    for _, c := range s {
        cnt[c - 'a']++
    }
    for _, c := range t {
        if cnt[c - 'a'] == 0 {
            return false
        }
        cnt[c - 'a']--
    }
    return true
}