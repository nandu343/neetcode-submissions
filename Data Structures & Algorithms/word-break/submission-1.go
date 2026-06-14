func wordBreak(s string, words []string) bool {
	    memo := make(map[string]bool)
    var dfs func(curr string, start int) bool
    dfs = func (curr string, start int) bool {
        if curr == s {
            return true
        }

        if v, e := memo[curr]; e == true {
            return v
        }

        possible := false
        for _, v := range words {
            if start + len(v) <= len(s) && s[start:start + len(v)] == v {
                possible = possible || dfs(curr + v, start + len(v))
                memo[curr + v] = possible
            }
        }
        
        return possible
    }
    
    return dfs("", 0)
}
