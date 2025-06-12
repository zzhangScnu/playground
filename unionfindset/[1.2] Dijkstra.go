package unionfindset

import "container/heap"
import h "code.byted.org/zhanglihua.river/playground/heap"

type Dijkstra struct {
	adjacent [][]int
	distance []int
	start    int
}

func NewDijkstra(n int, adjacent [][]int, start int) *Dijkstra {
	distance := make([]int, n)
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
	minHeap := &h.GoMinHeap{}
	heap.Init(minHeap)
	for _, to := range d.adjacent[d.start] {
		minHeap.Push(to)
	}
	for minHeap.Len() > 0 {
		cur := minHeap.Pop().(int)
		for _, to := range d.adjacent[cur] {
			if d.distance[cur]+d.weight(cur, to) > d.distance[to] {
				continue
			}
			d.distance[to] = d.distance[cur] + d.weight(cur, to)
			minHeap.Push(to)
		}
	}
}
