package heap

type Item struct {
	Num int
	Cnt int
}

type MinHeap struct {
	data []*Item
}

func NewMinHeap() MinHeap {
	return MinHeap{}
}

func (h *MinHeap) HeapifyByShiftDown(items []*Item) {
	h.data = make([]*Item, len(items))
	copy(h.data, items)
	for i := len(items) / 2; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *MinHeap) HeapifyByShiftUp(items []*Item) {
	for _, item := range items {
		h.Insert(item.Num, item.Cnt)
	}
}

func (h *MinHeap) shiftUp(idx int) {
	for idx > 0 && h.data[(idx-1)/2].Cnt > h.data[idx].Cnt {
		h.data[(idx-1)/2], h.data[idx] = h.data[idx], h.data[(idx-1)/2]
		idx = (idx - 1) / 2
	}
}

func (h *MinHeap) Insert(num, cnt int) {
	h.data = append(h.data, &Item{
		Num: num,
		Cnt: cnt,
	})
	h.shiftUp(h.Size() - 1)
}

func (h *MinHeap) ExtractMin() *Item {
	if h.Size() == 0 {
		return nil
	}
	val := h.PeekMin()
	h.data[0] = h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	h.shiftDown(0)
	return val
}

func (h *MinHeap) ReplaceMin(num, cnt int) {
	if h.Size() == 0 {
		h.Insert(num, cnt)
		return
	}
	h.data[0] = &Item{
		Num: num,
		Cnt: cnt,
	}
	h.shiftDown(0)
}

func (h *MinHeap) shiftDown(idx int) {
	MinPos := idx
	for {
		if idx*2+1 < h.Size() && h.data[idx*2+1].Cnt < h.data[MinPos].Cnt {
			MinPos = idx*2 + 1
		}
		if idx*2+2 < h.Size() && h.data[idx*2+2].Cnt < h.data[MinPos].Cnt {
			MinPos = idx*2 + 2
		}
		if MinPos == idx {
			break
		}
		h.data[idx], h.data[MinPos] = h.data[MinPos], h.data[idx]
		idx = MinPos
	}
}

func (h *MinHeap) PeekMin() *Item {
	if h.IsEmpty() {
		return nil
	}
	return h.data[0]
}

func (h *MinHeap) Size() int {
	return len(h.data)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MinHeap) Sort() []*Item {
	var res []*Item
	for h.Size() != 0 {
		res = append(res, h.ExtractMin())
	}
	return res
}
