package heap

import "container/heap"

// MedianFinder 中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
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
type MedianFinder struct {
	bigger  *GoMinHeap
	smaller *GoMaxHeap
}

func MedianFinderConstructor() MedianFinder {
	bigger, smaller := &GoMinHeap{}, &GoMaxHeap{}
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

/**
思路：
因为数据范围会动态增加，所以如果将有序整数列表维护在数组中，会带来搬迁数据的额外开销。
中位数位于有序整数列表的中间位置，考虑用两个数据结构分别维护较大的部分和较小的部分。
用最大堆和最小堆能满足要求。

想象整体数据排列，是一个正三角形，上小下大；从中位数位置切开两半，分别放进两个堆中，
且因为要求中位数，即从中间位置取，所以：
- 上半部分较小的数，需要倒置维护，即用最大堆；
- 下半部分较大的数，正序维护即可，即用最小堆。

插入新的元素时：
将其加入长度较短的堆中。这样动态调整能使得两个堆的大小差不会大于1。
如果加入上半部分中，则需在增加元素、重新堆化后，从堆口弹出一个元素，加入下半部分中。
这样会使得两个堆仍满足堆的性质、满足元素大小及数量的相对关系要求。

获取中位数时：
情况1：两个堆的长度一致。此时分别从堆口取元素，求平均值；
情况2：其中一个堆的元素数量比另一个堆大1。此时从长度更长的堆中取堆口元素。

注意：
Go中的heap，在push / pop等操作时，不能直接操作heap本身，需要调用heap包的库函数实现，
即this.bigger.Push(num)是不对的，需要heap.Push(this.bigger, num)。
*/
