package dijkstra

import (
	"container/heap"
	"math"
)

type Node struct {
	node     int // 当前节点
	distance int // 从起点到当前节点的最短路径的权重和
}

type MinHeap []*Node

func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Dijkstra struct {
	adjacent [][]int // 邻接表，存储图信息
	distance []int   // 结果集，distance[i]表示从start到节点i的最短路径的权重和
	start    int     // 起始节点
}

func NewDijkstra(n int, adjacent [][]int, start int) *Dijkstra {
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt
	}
	distance[start] = 0
	return &Dijkstra{
		adjacent: adjacent,
		distance: distance,
		start:    start,
	}
}

/*
根据实际情况，节点间的权重计算方式可能不同
*/
func (d *Dijkstra) weight(x, y int) int {
	return x - y
}

func (d *Dijkstra) calculateDistance() {
	// 优先级队列，队口元素为当前【从start开始的权重和最小的最短路径】到达的节点，体现贪心思想
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	// 将起点推入优先级队列。节点到自身的距离为0
	minHeap.Push(&Node{
		node:     d.start,
		distance: 0,
	})
	for minHeap.Len() > 0 {
		cur := minHeap.Pop().(*Node)
		// 若结果集中[start, cur]的权重和已经比队列中[start, cur]的权重和更小，即可能这条路径已经被其他节点更新优化过了，当前路径不再适用。
		// 即与cur相邻的节点to，其最短路径[start, to]，不会因为当前队列取出的子路径[start, cur]更新为更优权重和
		if cur.distance > d.distance[cur.node] {
			continue
		}
		// 获取cur的相邻节点
		for _, to := range d.adjacent[cur.node] {
			// 尝试更新to的结果，即找到[start, to]的一条更优最短路径
			// 根据cur计算出的[start, to]的最短路径的权重和 weight = [start, cur]的权重和 + [cur, to]的权重和
			// 即 weight = 结果集中start到cur的最短路径权重和 + 动态计算的cur到to的距离
			// 如果 重新计算的weight 比 记录在案的distance[to] 更优，则更新distance，
			// 并将to和weight作为候选入列，在之后用于更新to的相邻节点的路径权重和。
			if d.distance[cur.node]+d.weight(cur.node, to) > d.distance[to] {
				continue
			}
			d.distance[to] = d.distance[cur.node] + d.weight(cur.node, to)
			minHeap.Push(&Node{
				node:     to,
				distance: d.distance[to],
			})
		}
	}
}
