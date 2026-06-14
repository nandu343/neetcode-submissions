func characterReplacement(s string, k int) int {
	var set [26]int
    l := 0
    maxFreq := 1
    maxwindow := 1
    for r := 0; r < len(s); r++ {
        if r == 0 {
            set[int(rune(s[r])-'A')]++
            continue
        }

        set[s[r]-'A']++
        if set[s[r]-'A'] > maxFreq {
            maxFreq = set[s[r]-'A']
        }
        
        
        for (r-l+1) - maxFreq > k {
            set[s[l]-'A']--
            l++
            top := -1
            for _, v := range set {
                if top < v {
                    top = v
                }
            }
            maxFreq = top
        }

        if r - l + 1 > maxwindow {
            maxwindow = r - l + 1
        }
    }
    return maxwindow
}
