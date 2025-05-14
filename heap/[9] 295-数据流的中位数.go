package heap

import "container/heap"

// 中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
//
// 例如 arr = [2,3,4] 的中位数是 3 。
// 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
//
// 实现 MedianFinder 类:
//
// MedianFinder() 初始化 MedianFinder 对象。
// void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
// double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10⁻⁵ 以内的答案将被接受。
//
// 示例 1：
//
// 输入
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// 输出
// [null, null, null, 1.5, null, 2.0]
//
// 解释
// MedianFinder medianFinder = new MedianFinder();
// medianFinder.addNum(1);    // arr = [1]
// medianFinder.addNum(2);    // arr = [1, 2]
// medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
// medianFinder.addNum(3);    // arr[1, 2, 3]
// medianFinder.findMedian(); // return 2.0
//
// 提示:
//
// -10⁵ <= num <= 10⁵
// 在调用 findMedian 之前，数据结构中至少有一个元素
// 最多 5 * 10⁴ 次调用 addNum 和 findMedian

type MinHeap []int

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	bigger  *MinHeap
	smaller *MaxHeap
}

func MedianFinderConstructor() MedianFinder {
	bigger, smaller := &MinHeap{}, &MaxHeap{}
	heap.Init(bigger)
	heap.Init(smaller)
	return MedianFinder{
		bigger:  bigger,
		smaller: smaller,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.smaller.Len() >= this.bigger.Len() {
		heap.Push(this.smaller, num)
		heap.Push(this.bigger, heap.Pop(this.smaller))
	} else {
		heap.Push(this.bigger, num)
		heap.Push(this.smaller, heap.Pop(this.bigger))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.smaller.Len() == this.bigger.Len() {
		return float64((*this.smaller)[0]+(*this.bigger)[0]) / 2.0
	} else if this.smaller.Len() > this.bigger.Len() {
		return float64((*this.smaller)[0])
	} else {
		return float64((*this.bigger)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
