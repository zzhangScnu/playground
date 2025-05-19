package heap

const nonExistedVal = -1

type MaxHeap struct {
	data []int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{[]int{}}
}

func (h *MaxHeap) HeapifyByShiftDown(nums []int) {
	h.data = make([]int, len(nums))
	copy(h.data, nums)
	/**
	只需要从索引值为n/2的节点开始处理，直到0。
	因为叶子节点左右子节点都为空，和自身比较意义不大。
	自顶向下的处理过程中，自身、父节点、兄弟节点都会参与大小比较并移动到正确的位置。
	*/
	for i := len(nums) / 2; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *MaxHeap) HeapifyByShiftUp(nums []int) {
	/**
	data := make([]int, len(nums))
	copy(data, nums)
	for i := h.Size() - 1; i >= 0; i-- {
		h.shiftUp(i)
	}
	这种堆化方式是有问题的，因为自底向上的处理过程中，只会比较自身和父节点的大小，而忽略了兄弟节点。
	存在兄弟节点才是最大值，但被跳过的情况。
	一次性给切片赋值再堆化，需要用ShiftDown的方式。

	而下方的实现，因为每次插入一个新元素就堆化一次，保证新元素入堆前，堆已经是最大堆了，
	只需要判断当前入堆的元素，是否能成为父节点即可。
	逐个将元素加入切片再堆化，才可以用ShiftUp的方式。
	*/
	for _, num := range nums {
		h.Insert(num)
	}
}

/*
*
这里用的是索引从0开始的存储方式。
所以本节点索引是0，左右孩子索引就是1和2，对应idx * 2 + 1和idx * 2 + 2。
那么从孩子的idx倒推父节点，则为(idx - 1) / 2。

如果索引从1开始，
本节点索引是1，左右孩子索引就是2和3，对应idx * 2和idx * 2 + 1。
那么从孩子的idx倒推父节点，则为idx / 2。
*/
func (h *MaxHeap) shiftUp(idx int) {
	for idx > 0 && h.data[(idx-1)/2] < h.data[idx] { // 比较大小的条件也提上来，可以减少遍历次数。因为若不满足此条件，说明已经完全堆化
		h.data[(idx-1)/2], h.data[idx] = h.data[idx], h.data[(idx-1)/2]
		idx = (idx - 1) / 2
	}
}

func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.shiftUp(h.Size() - 1)
}

func (h *MaxHeap) ExtractMax() int {
	if h.Size() == 0 {
		return 0
	}
	val := h.PeekMax()
	h.data[0] = h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	h.shiftDown(0)
	return val
}

func (h *MaxHeap) shiftDown(idx int) {
	maxPos := idx
	for {
		if idx*2+1 < h.Size() && h.data[idx*2+1] > h.data[maxPos] {
			maxPos = idx*2 + 1
		}
		if idx*2+2 < h.Size() && h.data[idx*2+2] > h.data[maxPos] {
			maxPos = idx*2 + 2
		}
		if maxPos == idx {
			break
		}
		h.data[idx], h.data[maxPos] = h.data[maxPos], h.data[idx]
		idx = maxPos
	}
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

func (h *MaxHeap) Sort() []int {
	var res []int
	for h.Size() != 0 {
		res = append(res, h.ExtractMax())
	}
	return res
}
