func countComponents(n int, edges [][]int) int {
    adjList := make([][]int, n)

	for _, v := range edges {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}
	count := 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if e := visited[i]; e == true {
			continue
		}
		queue := []int{i}
		visited[i] = true
		for len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]
			for _, v := range adjList[node] {
				if e := visited[v]; e == true {
					continue
				}
				queue = append(queue, v)
				visited[v] = true
			}
		}
		count++
	}
	return count
}
