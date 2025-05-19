package heap

type GoMaxHeap []int

func (h GoMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h GoMaxHeap) Len() int {
	return len(h)
}

func (h GoMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *GoMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *GoMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
