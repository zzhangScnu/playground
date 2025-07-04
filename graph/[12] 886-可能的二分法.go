package graph

// 给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。
//
// 给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和 bi的人归入同一组。当可以用
// 这种方法将所有人分进两组时，返回 true；否则返回 false。
//
// 示例 1：
//
// 输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
// 输出：true
// 解释：group1 [1,4], group2 [2,3]
//
// 示例 2：
//
// 输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
// 输出：false
//
// 示例 3：
//
// 输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
// 输出：false
//
// 提示：
//
// 1 <= n <= 2000
// 0 <= dislikes.length <= 10⁴
// dislikes[i].length == 2
// 1 <= dislikes[i][j] <= n
// ai < bi
// dislikes 中每一组都 不同
func possibleBipartition(n int, dislikes [][]int) bool {
	graph := buildUpGraph(n, dislikes)
	color, visited := make([]bool, n+1), make([]bool, n+1)
	isBipartite := true
	var traverse func(graph [][]int, vertex int)
	traverse = func(graph [][]int, vertex int) {
		visited[vertex] = true
		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				color[neighbor] = !color[vertex]
				traverse(graph, neighbor)
			} else if color[neighbor] == color[vertex] {
				isBipartite = false
				return
			}
		}
	}
	for v := 1; v <= n; v++ {
		if !visited[v] {
			traverse(graph, v)
		}
	}
	return isBipartite
}

func buildUpGraph(n int, dislikes [][]int) [][]int {
	graph := make([][]int, n+1)
	for _, dislike := range dislikes {
		v, w := dislike[0], dislike[1]
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}
	return graph
}

/**
思路：
先将dislikes数组，即需要分进两个组的元素作为节点构建图，再基于该图运行二分图判定。

注意：
构建的是邻接表，而不是邻接矩阵，
所以需要用：
v, w := dislike[0], dislike[1]
graph[v] = append(graph[v], w)
graph[w] = append(graph[w], v)
对每个vertex做相邻节点的映射和拉链。

而不是：
v, w := dislike[0], dislike[1]
graph[v] = graph[w]
graph[w] = graph[v]
其实这种写法也不是邻接矩阵，真实的邻接矩阵写法是graph[v][w] = 1，表示v -> w存在一条有向边。
*/
