func orangesRotting(grid [][]int) int {
	count := 0
	queue := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i,j})
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}

	if count == 0 {
		return 0
	}

	delta := [4][2]int{{1,0},{-1,0},{0,1},{0,-1}}
	minute := 0
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			for _, v := range delta {
				x := node[0] + v[0]
				y := node[1] + v[1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
					continue
				}
				if grid[x][y] == 0 || grid[x][y] == 2 {
					continue
				}
				count--
				grid[x][y] = 2
				queue = append(queue, [2]int{x,y})
			}
		}
		minute++
	}
	if count == 0 {
		return minute - 1
	}
	return -1
}
