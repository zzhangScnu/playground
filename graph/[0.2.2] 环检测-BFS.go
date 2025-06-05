package graph

func detectCircleBFS(graph [][]int) bool {
	n := len(graph)
	if n == 0 {
		return false
	}
	visited, inDegree := make([]bool, n), make([]int, n)
	var queue []int
	for _, tos := range graph {
		for _, to := range tos {
			inDegree[to]++
		}
	}
	for vertex, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, vertex)
		}
	}
	var count int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] {
			continue
		}
		visited[cur] = true
		count++
		for _, neighbor := range graph[cur] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return count == n
}
