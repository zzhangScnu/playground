package linklist

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
// 示例 1：
//
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
//
// 示例 2：
//
// 输入：head = [1,1,1,2,3]
// 输出：[2,3]
//
// 提示：
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	var cur, next *ListNode = dummyHead, nil
	var duplicateVal int
	for cur != nil {
		next = cur.Next
		if next != nil && next.Next != nil && next.Val == next.Next.Val {
			duplicateVal = next.Val
			for next != nil && next.Val == duplicateVal {
				next = next.Next
			}
			cur.Next = next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}
