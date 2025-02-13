package linklist

import "container/heap"

// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
//
// 示例 2：
//
// 输入：lists = []
// 输出：[]
//
// 示例 3：
//
// 输入：lists = [[]]
// 输出：[]
//
// 提示：
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() any {
	if h.Len() <= 0 {
		return nil
	}
	val := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, cur := range lists {
		if cur != nil {
			heap.Push(minHeap, cur)
		}
	}
	dummyHead := &ListNode{}
	cur := dummyHead
	for minHeap.Len() > 0 {
		minNode := heap.Pop(minHeap).(*ListNode)
		cur.Next = minNode
		cur = cur.Next
		if minNode.Next != nil {
			heap.Push(minHeap, minNode.Next)
		}
	}
	return dummyHead.Next
}

/**
有几个注意点：

container/heap接口的实现：
1. Swap：应交换节点，而不是简单交换Val；
2. Less：无需判断数组越界，底层实现保证入参索引合法；
3. Pop：底层实现会先将堆头、堆尾交换，调用Pop后再重新堆化。所以Pop的实现返回堆尾数据并缩容即可。

container/heap接口的使用：
1. heap.Init：在使用前要调用，目的是堆的初始化；
2. heap.Push：入堆并堆化。而不是直接调用minHeap的Pop；
3. heap.Pop：出堆并堆化。而不是直接调用heap.Pop()。
*/

/**
第一次实现：想将所有链表中的所有节点都入堆，再依次出堆组成新的链表；
但更好的思路是，将【非空】的链表头节点入堆，
再边出堆边将下一个非空节点入堆。
*/
