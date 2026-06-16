func validTree(n int, edges [][]int) bool {
    
	if n == 1 && len(edges) == 0 {
        return true
    }

	adjMatrix := make([][]int, n)

	for _,v := range edges {
		adjMatrix[v[0]] = append(adjMatrix[v[0]], v[1])
		adjMatrix[v[1]] = append(adjMatrix[v[1]], v[0])
	}

	inOrders := make([]int, n)
	count := 0;
	for _, v := range edges {
		inOrders[v[1]]++
		inOrders[v[0]]++
		count++
	}
	
	if count != n - 1 {
		return false
	}
	queue := make([]int, 0)
	for i, v := range inOrders {
		if v == 1 {
			queue = append(queue, i)
		}
	}

	ans := make([]int, 0)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, v := range adjMatrix[node] {
			inOrders[v]--
			if inOrders[v] == 1 {
				queue = append(queue, v)
			}
		}
		ans = append(ans, node)
	}

	if len(ans) != n {
		return false
	}
	return true
}
