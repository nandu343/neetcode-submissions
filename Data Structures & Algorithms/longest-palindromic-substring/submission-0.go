func longestPalindrome(s string) string {
    lg := 0
	rg := 0
	var dfs func (l int, r int) bool
	memo := make(map[[2]int]bool)
	dfs = func (l int, r int) bool {
		if v, e := memo[[2]int{l,r}]; e {
			return v
		}

		if l > r {
			return true
		}

		if l == r {
			memo[[2]int{l,r}] = true
			return true
		}

		isPal := false
		if s[l] == s[r] && dfs(l+1, r - 1) {
			if r - l > rg - lg {
				rg = r
				lg = l
			}
			isPal = true
		}

		dfs(l + 1, r)
		dfs(l, r - 1)
		memo[[2]int{l,r}] = isPal
		return isPal

	}
	dfs(0, len(s)-1)
	return s[lg:rg+1]
}