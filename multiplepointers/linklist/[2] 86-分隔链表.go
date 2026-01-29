package linklist

import "code.byted.org/zhanglihua.river/playground/linklist"

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
//
// 示例 1：
//
// 输入：head = [1,4,3,2,5,2], x = 3
// 输出：[1,2,2,4,3,5]
//
// 示例 2：
//
// 输入：head = [2,1], x = 2
// 输出：[1,2]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200
func partition(head *linklist.ListNode, x int) *linklist.ListNode {
	smallerDummyHead, biggerDummyHead := &linklist.ListNode{Next: head}, &linklist.ListNode{Next: head}
	s, b := smallerDummyHead, biggerDummyHead
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			s.Next = cur
			s = s.Next
		} else {
			b.Next = cur
			b = b.Next
		}
	}
	b.Next = nil
	s.Next = biggerDummyHead.Next
	return smallerDummyHead.Next
}

/**
一开始想在原链表上就地移动节点，各种指针把自己绕晕了……
这种做法比较直观，先分再合。
唯一要注意的点是，因为在分链表的时候，是没有处理节点的下一个指针的指向的，
所以当原链表遍历完成后，要将保存较大值的链表的末尾节点指向nil，
否则它可能指向的还是某个较小值的节点。
（较小值链表不需要的原因是，它的末尾直接指向较大值链表了

另一种方式是，每次for循环中将节点连接到s和b链表后，都断开next指针对原链表节点的指向。

for cur := head; cur != nil; cur = cur.Next
这样的遍历方式，不用显式在循环最后另 cur = cur.Next。
*/
