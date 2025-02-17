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

import . "code.byted.org/zhanglihua.river/playground/2. linklist"

func sortList(head *ListNode) *ListNode {
	return doSortList(head, nil)
}

func doSortList(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	mid := findMiddle(head, tail)
	node1, node2 := doSortList(head, mid), doSortList(mid, tail)
	return merge2Lists(node1, node2)
}

func findMiddle(head, tail *ListNode) *ListNode {
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
