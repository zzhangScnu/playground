package graph

import (
	"slices"
)

var res []int

func topologicalSort(graph [][]int) []int {
	n := len(graph)
	visited, res = make([]bool, n), make([]int, n)
	for v := 0; v < n; v++ {
		traverseInTopologicalSort(graph, v)
	}
	slices.Reverse(res)
	return res
}

func traverseInTopologicalSort(graph [][]int, vertex int) {
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
		traverseInTopologicalSort(graph, neighbor)
	}
	res = append(res, vertex)
	path[vertex] = false
}
