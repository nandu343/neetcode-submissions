func coinChange(coins []int, amount int) int {
    memo := make(map[int]int)
	// min := math.MaxInt
	var dfs func ( sum int) int


	dfs = func (sum int) int {
		

		if sum == amount {
			return 0
		}

		if v, e := memo[sum]; e == true {
			return v
		}

		if sum > amount {
			return math.MaxInt
		}

		currmin := math.MaxInt
		for i := 0; i < len(coins); i++ {
			curr := dfs( sum + coins[i]) 
			if curr + 1 < currmin && curr != math.MaxInt {
				currmin = curr + 1
			}
		}
		memo[sum] = currmin
		return currmin
	}
	ans := dfs( 0)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
