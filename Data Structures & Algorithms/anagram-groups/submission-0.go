func groupAnagrams(strs []string) [][]string {
	sets := make(map[[26]int][]string)
    for _, v := range strs {
        var set [26]int
        for _, r := range v {
            set[int(r - 'a')]++
        }
        sets[set] = append(sets[set], v)
    }
    ans := make([][]string, len(sets))
    count := 0
    for _, v := range sets {
        ans[count] = v
        count++
    }
    return ans
}
