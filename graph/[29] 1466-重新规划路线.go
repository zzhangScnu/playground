package graph

// n 座城市，从 0 到 n-1 编号，其间共有 n-1 条路线。因此，要想在两座不同城市之间旅行只有唯一一条路线可供选择（路线网形成一颗树）。去年，交通运输
// 部决定重新规划路线，以改变交通拥堵的状况。
//
// 路线用 connections 表示，其中 connections[i] = [a, b] 表示从城市 a 到 b 的一条有向路线。
//
// 今年，城市 0 将会举办一场大型比赛，很多游客都想前往城市 0 。
//
// 请你帮助重新规划路线方向，使每个城市都可以访问城市 0 。返回需要变更方向的最小路线数。
//
// 题目数据 保证 每个城市在重新规划路线方向后都能到达城市 0 。
//
// 示例 1：
//
// 输入：n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
// 输出：3
// 解释：更改以红色显示的路线的方向，使每个城市都可以到达城市 0 。
//
// 示例 2：
//
// 输入：n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
// 输出：2
// 解释：更改以红色显示的路线的方向，使每个城市都可以到达城市 0 。
//
// 示例 3：
//
// 输入：n = 3, connections = [[1,0],[2,0]]
// 输出：0
//
// 提示：
//
// 2 <= n <= 5 * 10^4
// connections.length == n-1
// connections[i].length == 2
// 0 <= connections[i][0], connections[i][1] <= n-1
// connections[i][0] != connections[i][1]
func minReorder(n int, connections [][]int) int {
	graph := make([][]int, n)
	edges := make(map[[2]int]bool)
	for _, connection := range connections {
		from, to := connection[0], connection[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
		edges[[2]int{from, to}] = true
	}
	visited := make([]bool, n)
	var count int
	var traverse func(cur int)
	traverse = func(cur int) {
		visited[cur] = true
		for _, next := range graph[cur] {
			if visited[next] {
				continue
			}
			if !edges[[2]int{next, cur}] {
				count++
			}
			traverse(next)
		}
	}
	traverse(0)
	return count
}

/**
思路：
所有节点都能到达0 -> 从0出发能到达所有节点
使用递推的思路，若 0 能到达节点A，则 0 就能到达节点 A 的邻接节点 B、C……

所以构建两个数据结构：
- graph：双向(无向)图，为了能够触达所有可能的城市；
- edges：图中存在的有向边。存储边的原始方形，为了判断是否需要修改方向。(A -> B) 表示为edges[[2]int{A, B}] = true。

如果站在节点 A 上，不存在一条边能够从 A 到达 B，则需要对统计结果++。

注意，将 visited 的判断从 traverse 入口改为 for 循环中，只有未访问过的邻接节点才会进入下一层递归。
会降低无用的递归成本。
*/
