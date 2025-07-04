package kruskal

import (
	"code.byted.org/zhanglihua.river/playground/unionfindset"
	"math"
	"sort"
)

// 给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。
//
// 连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示
// val 的绝对值。
//
// 请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。
//
// 示例 1：
//
// 输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
// 输出：20
// 解释：
//
// 我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。
// 注意到任意两个点之间只有唯一一条路径互相到达。
//
// 示例 2：
//
// 输入：points = [[3,12],[-2,5],[-4,1]]
// 输出：18
//
// 示例 3：
//
// 输入：points = [[0,0],[1,1],[1,0],[-1,1]]
// 输出：4
//
// 示例 4：
//
// 输入：points = [[-1000000,-1000000],[1000000,1000000]]
// 输出：4000000
//
// 示例 5：
//
// 输入：points = [[0,0]]
// 输出：0
//
// 提示：
//
// 1 <= points.length <= 1000
// -10⁶ <= xi, yi <= 10⁶
// 所有点 (xi, yi) 两两不同。
func minCostConnectPoints(points [][]int) int {
	var connections [][]int
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			connections = append(connections, []int{i, j, calculateCost(points, i, j)})
		}
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	var cost int
	unionFindSet := unionfindset.NewUnionFindSet(len(points))
	for _, connection := range connections {
		if unionFindSet.IsConnected(connection[0], connection[1]) {
			continue
		}
		unionFindSet.Union(connection[0], connection[1])
		cost += connection[2]
	}
	return cost
}

func calculateCost(points [][]int, i, j int) int {
	a, b := points[i], points[j]
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

/**
套用最小生成树算法即可。
注意：
1. Kruskal原始算法维护的是一维形式的顶点坐标，而本题需通过二维坐标唯一确定一个顶点。
   所以通过points的索引下标i -> 顶点，做了一层映射转换，忽略了顶点的二维坐标信息；
2. 边的权重没有直接给出，需要通过节点提前进行计算并维护到图的存储信息中；
3. 并查集的初始化大小应为图中顶点的个数。最开始做的时候错误地初始化为len(connections)，实际应为len(points)。
*/
