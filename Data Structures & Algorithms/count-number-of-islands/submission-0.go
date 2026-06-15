func numIslands(grid [][]byte) int {
    count := 0
	delta := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	var bfs func (row int, col int)

	bfs = func (row int, col int) {
		grid[row][col] = '0'
		queue := make([][2]int, 0)
		queue = append(queue, [2]int{row, col})
		for len(queue) != 0 {
			k := len(queue)
			for i := 0; i < k; i++ {
				node := queue[0]
				queue = queue[1:]
				for j := 0; j < 4; j++ {
					newCords := [2]int{node[0] + delta[j][0], node[1] + delta[j][1]}
					if !isValid(newCords, grid) || grid[newCords[0]][newCords[1]] == '0' {
						continue
					}
					grid[newCords[0]][newCords[1]] = '0'
					queue = append(queue, newCords)
				}
			}
		}
		count++
	}
	
	for i, v := range grid {
		for j, v2 := range v {
			if v2 == '1' {
				bfs(i,j)
			}
		}
	}
	return count
}

func isValid(coord [2]int, grid [][]byte) bool {
	if coord[0] < 0 || coord[0] >= len(grid) || coord[1] < 0 || coord[1] >= len(grid[0]) {
		return false
	}
	return true
}
