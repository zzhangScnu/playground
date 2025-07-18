package dijkstra

import (
	"container/heap"
	"math"
)

// 有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城
// 市 fromi 开始，以价格 pricei 抵达 toi。
//
// 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便
// 宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
// 示例 1：
//
// 输入:
// n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0,
// dst = 3, k = 1
// 输出: 700
// 解释: 城市航班图如上
// 从城市 0 到城市 3 经过最多 1 站的最佳路径用红色标记，费用为 100 + 600 = 700。
// 请注意，通过城市 [0, 1, 2, 3] 的路径更便宜，但无效，因为它经过了 2 站。
//
// 示例 2：
//
// 输入:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
// 输出: 200
// 解释:
// 城市航班图如上
// 从城市 0 到城市 2 经过最多 1 站的最佳路径标记为红色，费用为 100 + 100 = 200。
//
// 示例 3：
//
// 输入：n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
// 输出：500
// 解释：
// 城市航班图如上
// 从城市 0 到城市 2 不经过站点的最佳路径标记为红色，费用为 500。
//
// 提示：
//
// 1 <= n <= 100
// 0 <= flights.length <= (n * (n - 1) / 2)
// flights[i].length == 3
// 0 <= fromi, toi < n
// fromi != toi
// 1 <= pricei <= 10⁴
// 航班没有重复，且不存在自环
// 0 <= src, dst, k < n
// src != dst
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	k++
	graph := buildGraph(n, flights)
	distance, nodeVisitedNum := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt
		nodeVisitedNum[i] = math.MaxInt
	}
	distance[src], nodeVisitedNum[src] = 0, 0
	minHeap := &LimitedMinHeap{}
	heap.Init(minHeap)
	srcNode := &LimitedNode{
		node:           src,
		distance:       0,
		nodeVisitedNum: 0,
	}
	heap.Push(minHeap, srcNode)
	for minHeap.Len() > 0 {
		cur := heap.Pop(minHeap).(*LimitedNode)
		if cur.node == dst {
			return cur.distance
		}
		if cur.nodeVisitedNum == k {
			continue
		}
		for _, to := range graph[cur.node] {
			toNode, weight := to[0], to[1]
			srcCurToDistance := cur.distance + weight
			srcCurToNodeVisitedNum := cur.nodeVisitedNum + 1
			if srcCurToDistance < distance[toNode] {
				distance[toNode] = srcCurToDistance
				nodeVisitedNum[toNode] = srcCurToNodeVisitedNum
			}
			if srcCurToDistance > distance[toNode] && srcCurToNodeVisitedNum > nodeVisitedNum[toNode] {
				continue
			}
			heap.Push(minHeap, &LimitedNode{
				node:           toNode,
				distance:       srcCurToDistance,
				nodeVisitedNum: srcCurToNodeVisitedNum,
			})
		}
	}
	return -1
}

type LimitedNode struct {
	node           int // 当前节点
	distance       int // 从起点到当前节点的最短路径的权重和
	nodeVisitedNum int // 从起点到当前节点的经过的节点个数
}

type LimitedMinHeap []*LimitedNode

func (h LimitedMinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h LimitedMinHeap) Len() int {
	return len(h)
}

func (h LimitedMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *LimitedMinHeap) Push(x any) {
	*h = append(*h, x.(*LimitedNode))
}

func (h *LimitedMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func buildGraph(n int, flights [][]int) [][][]int {
	graph := make([][][]int, n)
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		graph[from] = append(graph[from], []int{to, price})
	}
	return graph
}

/**
思路：
在基础的Dijkstra算法之上，叠加对中转次数对限制。
题目要求最大中转次数为k，即对于A -> B -> C：
- 中转次数为1，经由B中转；
- 访问节点数量为2：从A出发，需经过B、C才能到达终点。

在最短路径Dijkstra算法中，在原本维护【当前节点坐标】和【起点 -> 当前节点的路径权重和】的结构体中，
额外维护【起点 -> 当前节点的访问节点数量】。

- 因为每次都会从队列中弹出最优路径的端点，基于该节点BFS扩展新的路径，
  所以所有的可能性都会被覆盖，且路径间因节点入列的先后顺序天然保证合法；
- 因节点扩展在路径扩展后会被丢弃，取而代之入列的是新的节点，
  所以节点间维护的数据、表达的路径天然隔离，互不影响。

注意点：
- 因为出发点指定为src，所以需将记录全局最小路径的distance数组全初始化为最大值，方便参与后续取小择优。
  这里一开始只将[1, n)初始化为最大值，遗漏坐标0的单元格，导致结果错误；
- 因为之前的题目中，只有路径和权重一个维度，所以完全以优先队列的出列顺序计算择优即可。
  仅需维护全局最小权重和的记录。
  但本题多了一个访问节点数量的维度，就会存在可能性：当前路径和 > 全局最小路径和，但访问节点数量 < 全局最小访问数量，
  有可能在后续节点的遍历中，在访问节点数量限制反超的情况下找到更优解。
  所以这种情况也需要将路径端点入列。
  只有当路径和和访问节点数量都不占优的情况下，才需要剪枝。
  所以需要额外维护全局最小节点访问次数的记录。
*/
