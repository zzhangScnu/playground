package linklist

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
//
// 示例 2：
//
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
//
// 提示：
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	begin, from, to := dummyHead, dummyHead.Next, dummyHead.Next
	for {
		for i := 0; i < k-1; i++ {
			if to != nil {
				to = to.Next
			}
		}
		if to == nil {
			break
		}
		next := to.Next
		reversedHead, reversedTail := reverseBetweenNode(from, to)
		begin.Next = reversedHead
		reversedTail.Next = next
		begin, from, to = reversedTail, next, next
	}
	return dummyHead.Next
}

// 前闭后闭
func reverseBetweenNode(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var pre, cur *ListNode = nil, head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
