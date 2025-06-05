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

/**
思路：
DFS：路径遍历，递归+回溯，检查是否有同一个节点在同一条路径中被重复访问；
- visited数组：标记节点是否已访问，避免重复访问导致死循环；
- path数组：记录当前路径上已访问的节点，即单次递归的堆栈中的节点。如果在触底回溯重复访问，则表示有环。
  当单次遍历路径已扩展到无下一个可用节点，则需要进行回溯，回到先前的节点，进行其他节点的选择和访问。

因为图可能是非联通图，所以需要对图中的每一个节点进行遍历，防止遗漏。
*/
