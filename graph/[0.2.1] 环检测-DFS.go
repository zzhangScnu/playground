package graph

var visited []bool
var path []bool
var hasCircle bool

func detectCircle(graph [][]int) bool {
	n := len(graph)
	visited, path = make([]bool, n), make([]bool, n)
	for i := 0; i < n; i++ {
		traverse(graph, i)
	}
	return hasCircle
}

func traverse(graph [][]int, vertex int) {
	if visited[vertex] {
		return
	}
	visited[vertex] = true
	if path[vertex] {
		hasCircle = true
		return
	}
	path[vertex] = true
	for _, neighbor := range graph[vertex] {
		traverse(graph, neighbor)
	}
	path[vertex] = false
}
