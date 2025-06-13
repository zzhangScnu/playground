package kruskal

import (
	"code.byted.org/zhanglihua.river/playground/unionfindset"
	"sort"
)

/**
想象一下你是个城市基建规划者,地图上有N座城市,它们按以1到N的次序编号。
给你一些可连接的选项connections,其中每个选项connections[i] = [city1, city2, cost]表示
将城市city1和城市city2连接所要的成本为cost(连接是双向的,也就是说城市city1和城市city2相连也同样意味着城市city2和城市city1相连)。
计算使得每对城市都连通的最小成本。如果根据已知条件无法完成该项任务,则请你返回-1。

输入:N=3, connections = [[1,2,5],[1,3,6],[2,3,11]]
输出:6
解释:
选出任意2条边都可以连接所有城市,我们从中选取成本最小的2条。
*/

func minimumCost(n int, connections [][]int) int {
	unionFindSet := unionfindset.NewUnionFindSet(n + 1)
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	var cost int
	for _, connection := range connections {
		if unionFindSet.IsConnected(connection[0], connection[1]) {
			continue
		}
		unionFindSet.Union(connection[0], connection[1])
		cost += connection[2]
	}
	if unionFindSet.Size() == 2 {
		return cost
	}
	return -1
}

/**
思路：
1. 将边按权重升序排列；
2. 一边判断候选边是否有必要加入联通分量中：如果边的两个节点已经联通，则无需重复添加，徒增成本；
   否则加入并查集并累加成本。
使用贪心思想，从权重较小的边的局部最优，推导出权重和最小的图的全局最优。
*/
