package tree

const nonExistedVal = -1

type MaxHeap struct {
	data []int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{[]int{}}
}

func (h *MaxHeap) Heapify(nums []int) {
	h.data = nums
	for range h.data {
		h.doHeapify()
	}
}

func (h *MaxHeap) doHeapify() {
	idx := h.Size() - 1
	for idx > 0 {
		if h.data[idx/2] < h.data[idx] {
			h.data[idx/2], h.data[idx] = h.data[idx], h.data[idx/2]
		}
		idx /= 2
	}
}

func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.doHeapify()
}

func (h *MaxHeap) ExtractMax() int {
	val := h.PeekMax()
	h.data[0] = h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	idx := 0
	for idx < h.Size() {
		left, right := idx*2+1, idx*2+2
		pos := left
		if h.data[left] < h.data[right] {
			pos = right
		}
		if h.data[idx] > h.data[pos] {
			break
		}
		h.data[idx], h.data[pos] = h.data[pos], h.data[idx]
		idx = pos
	}
	return val
}

func (h *MaxHeap) PeekMax() int {
	if h.IsEmpty() {
		return nonExistedVal
	}
	return h.data[0]
}

func (h *MaxHeap) Size() int {
	return len(h.data)
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}
