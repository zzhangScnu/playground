package linklist

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
//
// 示例 2：
//
// 输入：head = [0,1,2], k = 4
// 输出：[2,0,1]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10⁹
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	length, tail := 1, head
	for ; tail.Next != nil; tail = tail.Next {
		length++
	}
	k %= length // 将 k 对链表长度取模，避免需要循环处理轮转
	if k == 0 {
		return head
	}
	cur := head
	for i := 0; i < length-k-1; i++ { // length - k - 1 确保 cur 指向新头部的【前一个节点】，即新的链表尾部
		cur = cur.Next
	}
	newHead := cur.Next // cur.Next 为新的链表头部
	cur.Next = nil      // cur 为新的链表尾部
	tail.Next = head    // 链表尾部链接上原本的链表头部
	return newHead
}

/**
思路：
模拟数组轮转
1. 首先通过遍历链表，计算链表长度，定位链表末尾；
2. 如果【链表长度】是【轮转次数】的模，则轮转结束后的链表跟原链表是一模一样的。故此时无需任何操作；
3. 否则通过k定位指针，通过操作指针操作将后边的链表拼接到前边的链表前。
*/
