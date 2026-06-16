func exist(board [][]byte, word string) bool {
	delta := [][2]int{{1,0}, {-1,0}, {0,1}, {0, -1}}
	visited := make(map[[2]int]struct{})
	var dfs func (row int, col int, leng int) bool
	dfs = func (row int, col int, leng int) bool {
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
			return false
		}
		if _, e := visited[[2]int{row,col}]; e == true {
			return false
		}
		if leng == len(word) && string(board[row][col]) == string(word[len(word)-1]) {
			return true
		}
		if string(board[row][col]) != string(word[leng-1]) {
			return false
		}
		visited[[2]int{row,col}] = struct{}{}
		ans := false
		for i := 0; i < 4; i++ {
			ans = ans || dfs(row + delta[i][0], col + delta[i][1], leng + 1)
		}
		delete(visited, [2]int{row,col})
		return ans
	}

	for i, v := range board {
		for j, _ := range v {
			if dfs(i, j, 1) == true {
				return true
			}
		}
	}
	return false
}
