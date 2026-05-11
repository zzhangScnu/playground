package heap

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
	bigger  *MinHeap
	smaller *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		bigger:  NewMinHeap(),
		smaller: NewMaxHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) { // 最终都是smaller多一个
	if len(this.smaller.data) == len(this.bigger.data) {
		this.bigger.Insert(num)
		this.smaller.Insert(this.bigger.Pop())
	} else {
		this.smaller.Insert(num)
		this.bigger.Insert(this.smaller.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.bigger.data) + len(this.smaller.data)
	if n%2 == 1 {
		return float64(this.smaller.data[0])
	}
	return float64(this.bigger.data[0]+this.smaller.data[0]) / 2.0 // 这里要除以2.0，否则会取整，先加转float再除
}

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (m *MinHeap) Insert(num int) {
	m.data = append(m.data, num)
	m.ShiftUp()
}

func (m *MinHeap) ShiftUp() {
	index := len(m.data) - 1
	for {
		parentIndex := (index - 1) / 2
		if m.data[index] >= m.data[parentIndex] {
			break
		}
		m.data[index], m.data[parentIndex] = m.data[parentIndex], m.data[index]
		index = parentIndex
	}
}

func (m *MinHeap) Pop() int { // 大小写可见性
	val := m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	if len(m.data) > 0 {
		m.shiftDown()
	}
	return val
}

func (m *MinHeap) shiftDown() {
	index := 0
	for {
		minIndex, minVal := index, m.data[index]
		left, right := index*2+1, index*2+2
		if left < len(m.data) && minVal > m.data[left] {
			minIndex, minVal = left, m.data[left]
		}
		if right < len(m.data) && minVal > m.data[right] {
			minIndex, minVal = right, m.data[right]
		}
		if minIndex == index {
			break
		}
		m.data[index], m.data[minIndex] = m.data[minIndex], m.data[index]
		index = minIndex
	}
}

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (m *MaxHeap) Insert(num int) {
	m.data = append(m.data, num)
	m.shiftUp()
}

func (m *MaxHeap) shiftUp() {
	index := len(m.data) - 1
	for {
		parentIndex := (index - 1) / 2
		if m.data[index] <= m.data[parentIndex] {
			break
		}
		m.data[index], m.data[parentIndex] = m.data[parentIndex], m.data[index]
		index = parentIndex
	}
}

func (m *MaxHeap) Pop() int {
	val := m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	if len(m.data) > 0 {
		m.shiftDown()
	}
	return val
}

func (m *MaxHeap) shiftDown() {
	index := 0
	for {
		maxIndex, maxVal := index, m.data[index]
		left, right := index*2+1, index*2+2
		if left < len(m.data) && maxVal < m.data[left] {
			maxIndex, maxVal = left, m.data[left]
		}
		if right < len(m.data) && maxVal < m.data[right] {
			maxIndex, maxVal = right, m.data[right]
		}
		if maxIndex == index {
			break
		}
		m.data[index], m.data[maxIndex] = m.data[maxIndex], m.data[index]
		index = maxIndex
	}
}

/**
之前简单地以为堆插入后，data 数组是从小到大排好序的，仅需取中间位置元素返回即可。
但实际上最小堆只承诺堆顶元素最小，实际上整个数组是乱序的。
*/
