func islandsAndTreasure(grid [][]int) {
    queue := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				queue = append(queue, [2]int{i,j})
			}
		}
	}

	count := 0
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			for _, v := range neighbors(node[0],node[1]) {
				if v[0] < 0 || v[0] >= len(grid) || v[1] < 0 || v[1] >= len(grid[0]) {
					continue
				}
				if grid[v[0]][v[1]] == -1 || grid[v[0]][v[1]] == 0 {
					continue
				}
				if count + 1 < grid[v[0]][v[1]] {
					grid[v[0]][v[1]] = count + 1
					queue = append(queue, [2]int{v[0],v[1]})
				}
			}
		}
		count++
	}

}

func neighbors (i int , j int) [][2]int {
	delta := [4][2]int{{1,0},{-1,0},{0,1},{0,-1}}
	ans := make([][2]int, 4)
	for a := 0; a < 4; a++ {
		ans[a] = [2]int{i+delta[a][0], j+delta[a][1]}
	}
	return ans
}
