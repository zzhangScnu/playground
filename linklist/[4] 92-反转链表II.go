package linklist

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	begin := dummyHead
	var pre, cur, next *ListNode
	for i := 1; i < left; i++ {
		begin = begin.Next
	}
	pre = begin.Next
	cur = pre.Next
	for i := left; i < right; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	begin.Next.Next = cur
	begin.Next = pre
	return dummyHead.Next
}

/**
1. 设置一个虚拟头节点，兼容以下情况：
- 空链表：如何避免访问越界情况；
- 从头开始反转链表：反转区间的前一个节点，如何找到此时是nil的情况。
2. 固定begin，反转区间的前一个节点。用以处理【区间反转后的头&尾】跟【原链表】重新连接；
3. 最后return dummyHead.Next而不是head：
- head：此时可能已经被反转，不再是头节点；
- dummyHead.Next：dummyHead的Next指针始终指向操作后的头节点：
	- 从中间开始反转：dummyHead.Next一直不变，就是链表的头节点；
	- 从头开始反转：begin即dummyHead，且begin.Next = pre，pre变成链表的头节点。
*/
