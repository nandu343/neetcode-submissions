func maxProfit(prices []int) int {
	l := 0
    r := 1
    maxCost := 0
    currCost := 0

    for r < len(prices) {
        if prices[l] > prices[r] {
            l = r
            r++
            continue
        }
        currCost = prices[r] - prices[l]
        if currCost > maxCost {
            maxCost = currCost
        }
        r++
    }
    return maxCost
}
