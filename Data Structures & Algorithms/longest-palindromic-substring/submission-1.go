func longestPalindrome(s string) string {
    n := len(s)
    if n == 0 { return "" }
    
    // Use a 2D slice for memoization. 
    // 0 = unvisited, 1 = false, 2 = true
    memo := make([][]int8, n)
    for i := range memo {
        memo[i] = make([]int8, n)
    }

    start, maxLen := 0, 1

    var isPalindrome func(l, r int) bool
    isPalindrome = func(l, r int) bool {
        if l >= r { return true }
        if memo[l][r] != 0 { return memo[l][r] == 2 }

        res := (s[l] == s[r]) && isPalindrome(l+1, r-1)
        
        if res {
            memo[l][r] = 2
            // Update global boundaries if this is the longest
            if (r - l + 1) > maxLen {
                maxLen = r - l + 1
                start = l
            }
        } else {
            memo[l][r] = 1
        }
        return res
    }

    // Trigger the check for all possible substrings
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            isPalindrome(i, j)
        }
    }

    return s[start : start+maxLen]
}