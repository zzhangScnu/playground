package graph

// 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连
// 。
//
// 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
//
// 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而
// isConnected[i][j] = 0 表示二者不直接相连。
//
// 返回矩阵中 省份 的数量。
//
// 示例 1：
//
// 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
// 输出：2
//
// 示例 2：
//
// 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
// 输出：3
//
// 提示：
//
// 1 <= n <= 200
// n == isConnected.length
// n == isConnected[i].length
// isConnected[i][j] 为 1 或 0
// isConnected[i][i] == 1
// isConnected[i][j] == isConnected[j][i]
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var traverse func(cur int)
	traverse = func(cur int) {
		visited[cur] = true
		for next, connectedFlag := range isConnected[cur] {
			if connectedFlag == 1 && !visited[next] {
				traverse(next)
			}
		}
	}
	var count int
	for i := 0; i < n; i++ {
		if !visited[i] {
			traverse(i)
			count++
		}
	}
	return count
}

/**
思路：
计算图的联通性，即图的联通分量。
采取 DFS 遍历的方式，限制重复访问，同时对每个联通图进行统计。

需要区分开邻接矩阵和原始图的区别。
只有图才需要通过[][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}的步长去做遍历。
邻接矩阵可以直接通过"跨越"的方式到达下一个节点。
*/
