func maxAreaOfIsland(grid [][]int) int {
	
	var neigh func (a int, b int) [][2]int
	neigh = func (a int, b int) [][2]int {
		nei := make([][2]int, 0, 4)
		delta := [][2]int{{1,0}, {-1,0}, {0, 1}, {0, -1}}
		for i := 0; i < 4; i++ {
			a2 := a + delta[i][0]
			b2 := b + delta[i][1]
			if a2 >= 0 && a2 < len(grid) && b2 >= 0 && b2 < len(grid[0]) {
				nei = append(nei, [2]int{a2, b2})
			}
		}
		return nei
	}

	var dfs func(i int, j int) int 
	dfs = func(i int, j int) int {
		if grid[i][j] == 0 {
			return 0
		}
		
		grid[i][j] = 0
		sum := 1
		neighbors := neigh(i, j)
		for _, v := range neighbors {
			sum += dfs(v[0], v[1])
		}
		return sum
	}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}


func max (a int, b int) int {
	if a > b {
		return a
	}
	return b
}