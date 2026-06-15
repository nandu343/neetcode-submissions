func canFinish(numCourses int, prerequisites [][]int) bool {
    
	inDegree := make([]int, numCourses)
	for _, v := range prerequisites {
		inDegree[v[1]]++
	}

	queue := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	ans := make([]int, 0)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, v := range prerequisites {
			if v[0] != node {
				continue
			}
			inDegree[v[1]]--
			if inDegree[v[1]] == 0 {
				queue = append(queue, v[1])
			}
		}
		ans = append(ans, node)
	}

	if len(ans) != numCourses {
		return false
	}
	return true
}


