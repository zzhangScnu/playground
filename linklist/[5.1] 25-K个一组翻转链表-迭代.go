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
		for i := 0; i < k-1 && to != nil; i++ {
			to = to.Next
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
	var pre, cur, next *ListNode = nil, head, nil
	for pre != tail {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}

/**
这种实现方式，本质就是模拟：
1. 在主方法中，找到四个关键节点：
- 需要反转的区间起点和终点，前闭后闭；
- 该区间的前驱和后驱，为了将反转后的子链表跟原链表连接起来。
2. to走k-1步，因为from占了一个位置，to再数k-1，才能凑齐k个节点；
3. 需要在每次反转前，提前把next存下来，防止操作后导致的断链。

注意：
- 一般迭代做法里的区间反转入参都是左闭右闭，返回反转后的头节点和尾节点，主要有两个作用：
	- 方便在原链表中进行连接；
	- 需要将下一个反转区间定位出来，如begin, from, to = reversedTail, next, next。
- to 的作用就是用于解决不足 k 个节点时需要保持原本顺序不变的问题，用 to 来标识需反转链表的尾部。
  所以循环的条件是 i < k-1 && to != nil，即当 to 为 nil 时，说明剩余链表长度已不足 k。
*/
