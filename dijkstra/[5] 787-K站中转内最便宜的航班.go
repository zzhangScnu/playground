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
