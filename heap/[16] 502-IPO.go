package heap

import "container/heap"

// 假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限
// ，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。
//
// 给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
//
// 最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
//
// 总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。
//
// 答案保证在 32 位有符号整数范围内。
//
// 示例 1：
//
// 输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
// 输出：4
// 解释：
// 由于你的初始资本为 0，你仅可以从 0 号项目开始。
// 在完成后，你将获得 1 的利润，你的总资本将变为 1。
// 此时你可以选择开始 1 号或 2 号项目。
// 由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
// 因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
//
// 示例 2：
//
// 输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
// 输出：6
//
// 提示：
//
// 1 <= k <= 10⁵
// 0 <= w <= 10⁹
// n == profits.length
// n == capital.length
// 1 <= n <= 10⁵
// 0 <= profits[i] <= 10⁴
// 0 <= capital[i] <= 10⁹
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minCapitals, maxProfits := &MinCapitalHeap{}, &GoMaxHeap{}
	heap.Init(minCapitals)
	heap.Init(maxProfits)
	n := len(profits)
	for i := 0; i < n; i++ {
		heap.Push(minCapitals, &Pair{capital[i], profits[i]})
	}
	for i := 0; i < k; i++ { // 最多可以运作 k 个资本
		for minCapitals.Len() > 0 && (*minCapitals)[0].capital <= w {
			pair := heap.Pop(minCapitals).(*Pair)
			heap.Push(maxProfits, pair.profit)
		}
		if maxProfits.Len() == 0 { // 如果已经没有足够 w 来支撑下一轮的 profits 选取，则结束循环
			break
		}
		w += heap.Pop(maxProfits).(int) // 注意这道题，profit 会加入 w 中，而对应的 capital 并不会从 w 中扣减
	}
	return w
}

type Pair struct {
	capital int
	profit  int
}

type MinCapitalHeap []*Pair

func (m MinCapitalHeap) Len() int {
	return len(m)
}

func (m MinCapitalHeap) Less(i, j int) bool {
	return m[i].capital < m[j].capital
}

func (m MinCapitalHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinCapitalHeap) Push(x any) {
	*m = append(*m, x.(*Pair))
}

func (m *MinCapitalHeap) Pop() any {
	n := m.Len()
	res := (*m)[n-1]
	*m = (*m)[:n-1]
	return res
}

/**
思路：
维护两个堆：
最小堆：以capital升序排序
最大堆：以profit降序排序

每次都从最小堆中获取最小的 capital，若手上资本w <= capital，则将 profit 加入最大堆中。
这种操作会维护【在实力允许条件下，可获取的最佳利润】序列。

若该序列不为空，每次从该序列中获取一个最佳利润，累加到资本w中。
更新资本 w 后，再去最小堆中循环 review 是否有可选择的 profit，若有则将其加入最大堆中。
重复此操作，直至用完所有机会k。

minCapital 堆的筛选逻辑：
minCapital 是按【启动资本】排序的小顶堆，堆顶始终是当前未处理项目中成本最低的。
当我们弹出堆顶项目时，说明它的成本 ≤ 当前资本 w。
而由于后续资本 w 会不断增加（因为利润为正），未来的资本一定 ≥ 现在的 w，所以这个项目未来也一定能被负担。
因此，只需将它加入 maxProfit 堆一次，就可以在后续所有轮次中被考虑，无需再放回 minCapital 堆。
*/

// 手撸堆版本
func findMaximizedCapitalII(k int, w int, profits []int, capital []int) int {
	capitalMinHeap0, profitMaxHeap0 := NewMinHeap0(), NewMaxHeap0()
	n := len(profits)
	for i := 0; i < n; i++ {
		capitalMinHeap0.Insert(capital[i], profits[i])
	}
	for i := 0; i < k; i++ {
		for capitalMinHeap0.Top() != nil && w >= capitalMinHeap0.Top()[0] {
			minCapital := capitalMinHeap0.Pop()
			profitMaxHeap0.Insert(minCapital[0], minCapital[1])
		}
		if len(profitMaxHeap0.data) == 0 {
			break
		}
		w += profitMaxHeap0.Pop()[1]
	}
	return w
}

type MinHeap0 struct {
	data [][]int
}

func NewMinHeap0() *MinHeap0 {
	return &MinHeap0{}
}

func (m *MinHeap0) Insert(capital int, profit int) {
	m.data = append(m.data, []int{capital, profit})
	m.shiftUp(len(m.data) - 1)
}

func (m *MinHeap0) shiftUp(index int) {
	for {
		parentIndex := (index - 1) / 2
		if m.data[index][0] >= m.data[parentIndex][0] {
			return
		}
		m.data[index], m.data[parentIndex] = m.data[parentIndex], m.data[index]
		index = parentIndex
	}
}

func (m *MinHeap0) Pop() []int {
	if len(m.data) == 0 {
		return nil
	}
	val := m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[0 : len(m.data)-1]
	if len(m.data) > 0 {
		m.shiftDown(0)
	}
	return val
}

func (m *MinHeap0) Top() []int {
	if len(m.data) == 0 {
		return nil
	}
	val := m.data[0]
	return val
}

func (m *MinHeap0) shiftDown(index int) {
	for {
		maxIndex, maxVal := index, m.data[index][0]
		if index*2+1 < len(m.data) && maxVal > m.data[index*2+1][0] {
			maxIndex, maxVal = index*2+1, m.data[index*2+1][0]
		}
		if index*2+2 < len(m.data) && maxVal > m.data[index*2+2][0] {
			maxIndex, maxVal = index*2+2, m.data[index*2+2][0]
		}
		if maxIndex == index {
			return
		}
		m.data[index], m.data[maxIndex] = m.data[maxIndex], m.data[index]
		index = maxIndex
	}
}

type MaxHeap0 struct {
	data [][]int
}

func NewMaxHeap0() *MaxHeap0 {
	return &MaxHeap0{}
}

func (m *MaxHeap0) Insert(capital int, profit int) {
	m.data = append(m.data, []int{capital, profit})
	m.shiftUp(len(m.data) - 1)
}

func (m *MaxHeap0) shiftUp(index int) {
	for {
		parentIndex := (index - 1) / 2
		if m.data[index][1] <= m.data[parentIndex][1] {
			return
		}
		m.data[index], m.data[parentIndex] = m.data[parentIndex], m.data[index]
		index = parentIndex
	}
}

func (m *MaxHeap0) Pop() []int {
	if len(m.data) == 0 {
		return nil
	}
	val := m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[0 : len(m.data)-1]
	if len(m.data) > 0 { // 注意这里的判断，需要弹出后堆内仍有元素，才触发重排
		m.shiftDown(0)
	}
	return val
}

func (m *MaxHeap0) shiftDown(index int) {
	for {
		minIndex, minVal := index, m.data[index][1]
		if index*2+1 < len(m.data) && minVal < m.data[index*2+1][1] {
			minIndex, minVal = index*2+1, m.data[index*2+1][1]
		}
		if index*2+2 < len(m.data) && minVal < m.data[index*2+2][1] {
			minIndex, minVal = index*2+2, m.data[index*2+2][1]
		}
		if minIndex == index {
			return
		}
		m.data[index], m.data[minIndex] = m.data[minIndex], m.data[index]
		index = minIndex
	}
}
