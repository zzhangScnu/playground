package kruskal

import (
	"container/heap"
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

type Edge struct {
	To       int
	Distance int
}

type EdgeMinHeap []*Edge

func (p EdgeMinHeap) Len() int {
	return len(p)
}

func (p EdgeMinHeap) Less(i, j int) bool {
	return p[i].Distance < p[j].Distance
}

func (p *EdgeMinHeap) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *EdgeMinHeap) Push(x any) {
	*p = append(*p, x.(*Edge))
}

func (p *EdgeMinHeap) Pop() any {
	n := p.Len()
	val := (*p)[n-1]
	*p = (*p)[:n-1]
	return val
}

type Prim struct {
	minHeap     *EdgeMinHeap
	visited     []bool
	totalWeight int
}

func NewPrim(n int) *Prim {
	minHeap := &EdgeMinHeap{}
	heap.Init(minHeap)
	return &Prim{
		minHeap:     minHeap,
		visited:     make([]bool, n),
		totalWeight: 0,
	}
}

func (p *Prim) IsAllConnected() bool {
	for _, flag := range p.visited {
		if !flag {
			return false
		}
	}
	return true
}

func (p *Prim) GetTotalWeight() int {
	return p.totalWeight
}

func (p *Prim) GenerateMinTree(graph [][][]int) {
	start := 0
	heap.Push(p.minHeap, &Edge{
		To:       start,
		Distance: 0,
	})
	for p.minHeap.Len() > 0 {
		cur := heap.Pop(p.minHeap).(*Edge)
		if p.visited[cur.To] {
			continue
		}
		p.visited[cur.To] = true
		p.totalWeight += int(cur.Distance)
		for _, neighbor := range graph[cur.To] {
			if p.visited[neighbor[0]] {
				continue
			}
			heap.Push(p.minHeap, &Edge{
				To:       neighbor[0],
				Distance: neighbor[1],
			})
		}
	}
}

func minimumCost(n int, connections [][]int) int {
	var buildGraph func(n int, connections [][]int) [][][]int
	buildGraph = func(n int, connections [][]int) [][][]int {
		graph := make([][][]int, n)
		for _, connection := range connections {
			v, w, cost := connection[0]-1, connection[1]-1, connection[2]
			graph[v] = append(graph[v], []int{w, cost})
			graph[w] = append(graph[w], []int{v, cost})
		}
		return graph
	}
	graph := buildGraph(n, connections)
	prim := NewPrim(n)
	prim.GenerateMinTree(graph)
	if !prim.IsAllConnected() {
		return -1
	}
	return prim.GetTotalWeight()
}

/**
可直接套用Prim算法解决。
注意题目中给出节点范围[1, N]，故在构造邻接表时需减去1。
*/
