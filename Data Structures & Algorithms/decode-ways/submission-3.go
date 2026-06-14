func numDecodings(digits string) int {
	

    memo := make(map[string]int)

    var dfs func (digi string) int 
    dfs = func (digi string) int {
        
        if len(digi) == 0 {
            return 1
        }
        if digi[0] == '0' {
            return 0
        }

        
        if len(digi) == 2 || len(digi) == 1 {
            dig, _ := strconv.Atoi(digi)
            if dig < 26 && dig > 10 && dig != 20{
                return 2
            } else if dig < 11 || dig == 20 {
                return 1
            }
        }


        if v, e := memo[digi]; e == true {
            return v
        }

        count := 0

        if v, _ := strconv.Atoi(digi[:2]); v > 26 {
            count = dfs(digi[1:])
        } else {
            count = dfs(digi[1:]) + dfs(digi[2:])
        }

        memo[digi] = count

        return count
        
    }

    return dfs(digits)
}
