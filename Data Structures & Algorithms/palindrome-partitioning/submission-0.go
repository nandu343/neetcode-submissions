func partition(s string) [][]string {
	
	var check func (a string) bool
	check = func (a string) bool {
		if len(a) == 0 {
			return false
		}
		l := 0
		r := len(a) - 1
		for l < r {
			if a[l] != a[r] {
				return false
			}
			r--
			l++
		}
		return true
	}
	
	global := make([][]string, 0)
	var dfs func  (start int, palin []string)
	dfs = func (start int, palin []string) {
		if start == len(s) {
			temp := make([]string, len(palin))
			copy(temp, palin)
			global = append(global, temp)
			return
		}
		if start > len(s) {
			return
		}

		for i := start + 1; i < len(s) + 1; i++ {
			if check(s[start:i]) {
				dfs(i, append(palin, s[start:i]))
			}
		}
	}

	dfs(0, []string{})
	return global
}
