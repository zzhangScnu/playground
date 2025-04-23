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
func reverseKGroupRecursively(head *ListNode, k int) *ListNode {
	to := head
	for i := 0; i < k; i++ {
		if to == nil {
			return head
		}
		to = to.Next
	}
	reversedHead := reverseBetweenNodeRecursively(head, to)
	head.Next = reverseKGroupRecursively(to, k)
	return reversedHead
}

// 前闭后开
func reverseBetweenNodeRecursively(head *ListNode, tail *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/**
思路：
1. 反转以head为头节点的长度为k的子链表；
2. 将k+1个节点作为头节点，递归调用，反转剩余的子链表；
3. 将上述两个子链表连接起来。

递归的方式，相比起迭代，
1. 每次反转区间，都是左闭右开：
- 否则在主方法里，需要额外找到本次操作子链表的前驱，比较麻烦；
- 所以主方法的to，是需要走k步的；
- 所以子方法的结束条件，是cur == tail；
2. 主方法里，相比起迭代的前后都要连接上，这里只处理前驱的连接：
- 所以子方法只返回子链表反转后的头节点，且是pre，而不是tail。
*/
