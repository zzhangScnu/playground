package graph

// 存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。给你一个二维数组 graph ，其中 graph[u]
// 是一个节点数组，由节点 u 的邻接节点组成。形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有
// 以下属性：
//
// 不存在自环（graph[u] 不包含 u）。
// 不存在平行边（graph[u] 不包含重复值）。
// 如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
// 这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
//
// 二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称
// 为 二分图 。
//
// 如果图是二分图，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
// 输出：false
// 解释：不能将节点分割成两个独立的子集，以使每条边都连通一个子集中的一个节点与另一个子集中的一个节点。
//
// 示例 2：
//
// 输入：graph = [[1,3],[0,2],[1,3],[0,2]]
// 输出：true
// 解释：可以将节点分成两组: {0, 2} 和 {1, 3} 。
//
// 提示：
//
// graph.length == n
// 1 <= n <= 100
// 0 <= graph[u].length < n
// 0 <= graph[u][i] <= n - 1
// graph[u] 不会包含 u
// graph[u] 的所有值 互不相同
// 如果 graph[u] 包含 v，那么 graph[v] 也会包含 u
func isBipartite(graph [][]int) bool {
	n := len(graph)
	visited, color := make([]bool, n), make([]bool, n)
	isBipartite := true
	var traverse func(graph [][]int, vertex int)
	traverse = func(graph [][]int, vertex int) {
		visited[vertex] = true
		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				color[neighbor] = !color[vertex]
				traverse(graph, neighbor)
			} else if color[neighbor] != !color[vertex] {
				isBipartite = false
				return
			}
		}
	}
	for v := 0; v < n; v++ {
		if !visited[v] {
			traverse(graph, v)
		}
	}
	return isBipartite
}

/**
思路：
遍历图的过程中，用两种颜色给节点染色，同时判断相邻节点颜色需不同。
visited数组：标记已访问节点，避免重复访问导致死循环。visited[i]表示节点i是否已访问；
color数组：true/false表示不同的颜色。color[i]表示节点i染上的颜色；
理论上这两个数组是可以统一的，但为了语义清晰，保持独立维护。

注意，因为图不一定是联通图，所以需要对每一个节点都执行一次遍历，在入口处判断未访问才调用方法。

一开始用的是DFS遍历的通用套路去遍历图，即定义上下左右四个方向的movements数组来控制节点遍历方向，
但题目给的是一个邻接表，而非邻接矩阵，
所以通过for循环遍历graph[vertex]对应的数组，就可以访问vertex相邻的所有neighbor节点。
*/
