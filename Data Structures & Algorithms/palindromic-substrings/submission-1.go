func countSubstrings(s string) int {
	memo := make(map[[2]int][2]int)
	var dfs func (i int, j int) (int, int)
	dfs = func (i int, j int) (int, int) {
		if v, e := memo[[2]int{i,j}]; e {
			return v[0], v[1]
		}
		if i > j {
			return 0, 1
		}
		if i == j {
			memo[[2]int{i,j}] = [2]int{0,1}
			return 1, 1
		}

		count, e := dfs(i + 1, j - 1)
		if s[i] == s[j] && e == 1 {
			count += 1
		} else {
			e = -1
		}

		a, _ := dfs(i + 1, j)
		b, _ := dfs(i, j - 1)

		count += (a + b)
		memo[[2]int{i,j}] = [2]int{0,e}

		return count, e
	}
	ans, _ := dfs(0, len(s)-1)
	return ans
}
