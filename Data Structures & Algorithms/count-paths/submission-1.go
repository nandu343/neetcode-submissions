func uniquePaths(m int, n int) int {
    matrix := make([][]int, m)
	for i := range matrix {
    matrix[i] = make([]int, n)
}

for j := 0; j < n; j++ {
        matrix[0][j] = 1
    }

    // First column
    for i := 0; i < m; i++ {
        matrix[i][0] = 1
    }



	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum := 0
			sum += matrix[i - 1][j]
			sum += matrix[i][j-1]

			matrix[i][j] = sum

		}
	}
	return matrix[m-1][n-1]
}
