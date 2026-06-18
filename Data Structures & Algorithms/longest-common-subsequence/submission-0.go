func longestCommonSubsequence(text1 string, text2 string) int {
    var dfs func (l int, r int) int
	memo := make(map[[2]int]int)
	dfs = func (l int, r int) int {
		if l >= len(text1) || r >= len(text2) {
			return 0
		}
		
		if v, e := memo[[2]int{l,r}]; e == true {
			return v
		}

		if text1[l] == text2[r] {
			ans := dfs(l + 1, r + 1) + 1
			memo[[2]int{l,r}] = ans
			return ans
		}
		ans := max(dfs(l, r + 1), dfs(l + 1, r))
		memo[[2]int{l,r}] = ans
		return ans
	}
	return dfs(0,0)
}

func max (a int, b int) int {
	if a > b {
		return a
	}
	return b
}
