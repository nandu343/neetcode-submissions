func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]struct{})
    l := 0
    maxcount := 0
    count := 0
    for r := 0; r < len(s); r++ {
        if _, e := set[rune(s[r])]; e == true {
            for {
                check := s[l]
                delete(set, rune(s[l]))
                l++
                count--
                if rune(check) == rune(s[r]) {
                    break
                }
            }
        }
        set[rune(s[r])] = struct{}{}
        count++
        if count > maxcount {
            maxcount = count
        }
    }
    return maxcount
}
