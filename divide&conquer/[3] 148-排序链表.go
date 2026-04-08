package divide_conquer

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 示例 1：
//
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
//
// 示例 2：
//
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]
//
// 示例 3：
//
// 输入：head = []
// 输出：[]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

import . "code.byted.org/zhanglihua.river/playground/linklist"

func sortList(head *ListNode) *ListNode {
	return doSortList(head, nil)
}

func doSortList(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail { // 因为区间是左闭右开 [head, tail)，如果 head 和 tail 挨着，说明已将链表切分到单个节点了，可以开始 merge 操作
		head.Next = nil // 将节点与原链表断开，避免藕断丝连
		return head
	}
	mid := findMiddle(head, tail)
	node1, node2 := doSortList(head, mid), doSortList(mid, tail)
	return merge2Lists(node1, node2)
}

func findMiddle(head, tail *ListNode) *ListNode {
	slow, fast := head, head
	for fast != tail && fast.Next != tail { // 注意这里的终止条件是 tail，而不是 nil
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/**
跟23一样，先切分直至剩余单个节点，触底回溯时再合并链表。
注意当head.Next == tail时，需要将head.Next置为nil，否则在调用merge2Lists时会陷入无限循环。

时间复杂度：
宽度：合并次数。无论在哪一层，合并都需要处理 n 个元素，n
深度：切分次数。每次将链表劈成两半，直到剩 1 个元素，log(n)
时间复杂度 == O(nlog(n))
*/
