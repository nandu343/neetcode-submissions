func minCostClimbingStairs(cost []int) int {
    costs := []int{cost[0],cost[1]}
	for i := 2; i < len(cost); i++ {
		if costs[i - 1] < costs[i - 2] {
			costs = append(costs, cost[i] + costs[i - 1])
			fmt.Println("a: ", cost[i] + costs[i - 1])
		} else {
			costs = append(costs, cost[i] + costs[i -2 ])
			fmt.Println("b: ", cost[i] + costs[i - 2])
		}
	}
	return min(costs[len(costs) - 1], costs[len(costs) - 2])
}

func min (a int, b int) int {
	if a > b {
		return b
	}
	return a
}
