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
func allPathsSourceTargetByAdjacencyList(graph [][]int) [][]int {
	var path []int
	var result [][]int
	var traverse func(graph [][]int, from, to int)
	traverse = func(graph [][]int, from, to int) {
		if from == to {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for _, next := range graph[from] {
			path = append(path, next)
			traverse(graph, next, to)
			path = path[:len(path)-1]
		}
	}
	path = append(path, 0)
	traverse(graph, 0, len(graph)-1)
	return result
}

/**
邻接矩阵：Adjacency Matrix，用二维数组表示节点之间的可达性。空间复杂度为O(n^2)，适合存储稠密图；
邻接表：Adjacency List，用数组+链表表示节点之间的可达性。空间复杂度为O(n+m)，适合存储稀疏图。

寻找所有可达路径：
本题中的图是有向无环图，没有重复访问的可能，因此不需要额外的visited数组来记录是否访问过某个节点。

深度优先搜索：
递归+回溯
在for循环中选择不同的分岔路径，通过递归深入搜索。
注意，因为路径总是由0开始的，最初就应将0加入path中，再通过0向不同节点发散。
*/
