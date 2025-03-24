package graph

// 给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
//
// graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。
//
// 示例 1：
//
// 输入：graph = [[1,2],[3],[3],[]]
// 输出：[[0,1,3],[0,2,3]]
// 解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
//
// 示例 2：
//
// 输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
// 输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
//
// 提示：
//
// n == graph.length
// 2 <= n <= 15
// 0 <= graph[i][j] < n
// graph[i][j] != i（即不存在自环）
// graph[i] 中的所有元素 互不相同
// 保证输入为 有向无环图（DAG）

func allPathsSourceTargetByAdjacencyMatrix(graph [][]int) [][]int {
	var path []int
	var result [][]int
	matrix := constructGraph(graph)
	var traverse func(matrix [][]int, cur, des int)
	traverse = func(matrix [][]int, cur, des int) {
		if cur == des {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for next, reachable := range matrix[cur] {
			if reachable == 1 {
				path = append(path, next)
				traverse(matrix, next, des)
				path = path[:len(path)-1]
			}
		}
	}
	path = append(path, 0)
	traverse(matrix, 0, len(graph)-1)
	return result
}

func constructGraph(graph [][]int) [][]int {
	n := len(graph)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for from, tos := range graph {
		for _, to := range tos {
			matrix[from][to] = 1
		}
	}
	return matrix
}

/**
单向无环图没有重复访问的可能，所以不需要记录访问过的节点
*/
