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
	for i := 0; i < k; i++ {
		if minCapitals.Len() > 0 {
			pair := heap.Pop(minCapitals).(*Pair)
			if pair.capital <= w {
				heap.Push(maxProfits, pair.profit)
			}
		}
		if maxProfits.Len() == 0 {
			break
		}
		w += heap.Pop(maxProfits).(int)
	}
	return w
}

type Pair struct {
	capital int
	profit  int
}

type MinCapitalHeap []Pair

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
	*m = append(*m, x.(Pair))
}

func (m *MinCapitalHeap) Pop() any {
	n := m.Len()
	res := (*m)[n-1]
	*m = (*m)[:n-1]
	return res
}

/**
todo：
在这个算法中，从 minCapital 堆中弹出的项目不需要再加回去，核心原因是每个项目只能被选择一次，且算法的逻辑确保了我们不会错过任何可做的项目。
具体来说，有两个关键原因：
项目的不可逆性
每个项目一旦被弹出 minCapital 堆并加入 maxProfit 堆，就意味着它已经被纳入 “当前可做的项目池”。即使这次没被选中（比如本轮选择了利润更高的项目），但由于资本 w 只会增加不会减少（利润为正），后续轮次中，这个项目仍然属于 “可负担” 的范围，会一直留在 maxProfit 堆中等待被选择。
例如：假设当前资本是 10，minCapital 堆中有项目 A（成本 5，利润 2）和项目 B（成本 8，利润 5）。
第一轮：A 和 B 都会被弹出 minCapital，加入 maxProfit 堆，最终选择 B（利润更高），资本变为 15。
第二轮：maxProfit 堆中还剩 A，此时不需要再从 minCapital 中找 A（已经在 maxProfit 里了），直接选择即可。
minCapital 堆的筛选逻辑
minCapital 是按 “启动资本” 排序的小顶堆，堆顶始终是当前未处理项目中成本最低的。
当我们弹出堆顶项目时，说明它的成本 ≤ 当前资本 w。而由于后续资本 w 会不断增加（因为利润为正），未来的资本一定 ≥ 现在的 w，所以这个项目未来也一定能被负担。
因此，只需将它加入 maxProfit 堆一次，就可以在后续所有轮次中被考虑，无需再放回 minCapital 堆。
简单说：一旦项目进入 maxProfit 堆，就永久处于 “可被选择” 的状态，而 minCapital 堆的作用只是 “初次筛选” 出可负担的项目，后续无需重复处理。这种设计既保证了正确性，又减少了不必要的操作，提高了效率
*/
