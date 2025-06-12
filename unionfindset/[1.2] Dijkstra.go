package unionfindset

import (
	"container/heap"
	"math"
)

type Node struct {
	node     int
	distance int
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
	adjacent [][]int
	distance []int
	start    int
}

func NewDijkstra(n int, adjacent [][]int, start int) *Dijkstra {
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance = math.MaxInt
	}
	distance[start] = 0
	return &Dijkstra{
		adjacent: adjacent,
		distance: distance,
		start:    start,
	}
}

func (d *Dijkstra) weight(x, y int) int {
	return x - y
}

func (d *Dijkstra) calculateDistance() {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	minHeap.Push(&Node{
		node:     d.start,
		distance: 0,
	})
	for minHeap.Len() > 0 {
		cur := minHeap.Pop().(*Node)
		if cur.distance > d.distance[cur.node] {
			continue
		}
		for _, to := range d.adjacent[cur.node] {
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
