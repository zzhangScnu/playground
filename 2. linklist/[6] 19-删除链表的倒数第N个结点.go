package linklist

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
// 示例 2：
//
// 输入：head = [1], n = 1
// 输出：[]
//
// 示例 3：
//
// 输入：head = [1,2], n = 1
// 输出：[1]
//
// 提示：
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// 进阶：你能尝试使用一趟扫描实现吗？
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

/**
倒数第n个节点，即为正数第len-n+1个节点。
先让快指针数n个节点，再让快慢指针同时前进；
i := 0; i < n+1 不是因为虚拟，是因为让slow指向被删节点的上一个节点。
*/
