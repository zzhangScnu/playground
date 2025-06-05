package graph

func topologicalSortBFS(graph [][]int) []int {
	var res []int
	n := len(graph)
	if n == 0 {
		return res
	}
	inDegree := make([]int, n)
	var queue []int
	for _, tos := range graph {
		for _, to := range tos {
			inDegree[to]++
		}
	}
	for vertex, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, vertex)
			res = append(res, vertex)
		}
	}
	var count int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		for _, neighbor := range graph[cur] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
				res = append(res, neighbor)
			}
		}
	}
	return res
}
