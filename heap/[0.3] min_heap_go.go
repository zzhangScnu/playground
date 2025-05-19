package heap

type GoMinHeap []int

func (h GoMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h GoMinHeap) Len() int {
	return len(h)
}

func (h GoMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *GoMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *GoMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
