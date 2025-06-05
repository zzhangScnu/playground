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
		}
	}
	var count int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		count++
		for _, neighbor := range graph[cur] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if count != n {
		return []int{}
	}
	return res
}

/**
思路：
由于BFS层序扩张、优先处理入度为零的节点的特性，从队列中出列的顺序即为拓扑排序顺序。

注意，如果图中有环，则显然无法完成拓扑排序，此时应显式返回空结果集。
*/
