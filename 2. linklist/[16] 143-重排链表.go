package linklist

import "container/list"

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// # L0 → L1 → … → Ln - 1 → Ln
//
// 请将其重新排列后变为：
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1：
//
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
//
// 示例 2：
//
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]
//
// 提示：
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
func reorderList(head *ListNode) {
	p := head
	ls := list.New()
	for p != nil {
		ls.PushBack(p)
		p = p.Next
	}
	var count int
	newHead := &ListNode{}
	q := head
	var element *list.Element
	for ls.Len() > 0 {
		if count%2 == 0 {
			element = ls.Front()
			q.Next = element.Value.(*ListNode)
			ls.Remove(element)
		} else {
			element = ls.Back()
			q.Next = element.Value.(*ListNode)
			ls.Remove(element)
		}
		count++
		q = q.Next
	}
	q.Next = nil
	head = newHead

}
