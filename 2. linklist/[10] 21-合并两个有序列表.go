package linklist

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例 1：
//
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
// 示例 2：
//
// 输入：l1 = [], l2 = []
// 输出：[]
//
// 示例 3：
//
// 输入：l1 = [], l2 = [0]
// 输出：[0]
//
// 提示：
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p, q := list1, list2
	dummyHead := &ListNode{}
	cur := dummyHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}
	return dummyHead.Next
}

/**
双指针思路。链表题基本适用。
唯一要注意的点是，当较短链表遍历完成后，将长链表剩余部分拼接到结果列表尾部，
一开始是这么写的：
	for p != nil {
		cur.Next = p
		cur = cur.Next
		p = p.Next
	}
	for q != nil {
		cur.Next = q
		cur = cur.Next
		q = q.Next
	}
但，	cur = cur.Next / p = p.Next这两句不要也罢……
*/
