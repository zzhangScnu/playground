package linklist

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例 1：
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
//
// 示例 2：
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
//
// 示例 3：
//
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
// 提示：
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, q, r := l1, l2, res
	var sum, carry int
	for p != nil || q != nil || carry > 0 {
		sum = carry
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		carry = sum / 10
		r.Next = &ListNode{Val: sum % 10}
		r = r.Next
	}
	return res.Next
}

/**
思路：
不断向前累加，模拟进位。
其中：
进位 = sum / 10；
本位 = sum % 10。

用p != nil || q != nil || carry > 0作为循环条件且在其中分别判断p != nil和q != nil，再进行对应的加总和指针移动，
是比较简洁的做法。
*/

/**
之前的写法，比较冗余：
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, q, r := l1, l2, res
	var sum, carry int
	for p != nil && q != nil {
		sum = p.Val + q.Val + carry
		carry = sum / 10
		r.Next = &ListNode{Val: sum % 10}
		p, q, r = p.Next, q.Next, r.Next
	}
	for p != nil {
		sum = p.Val + carry
		carry = sum / 10
		r.Next = &ListNode{Val: sum % 10}
		p, r = p.Next, r.Next
	}
	for q != nil {
		sum = q.Val + carry
		carry = sum / 10
		r.Next = &ListNode{Val: sum % 10}
		q, r = q.Next, r.Next
	}
	if carry > 0 { // 还把这段逻辑漏掉了，导致丢失了最后的进位……
		r.Next = &ListNode{Val: carry}
	}
	return res.Next
}
*/

/**
思路：
不断向前累加，模拟进位。
其中：
进位 = sum / 10；
本位 = sum % 10。

用p != nil || q != nil || carry > 0作为循环条件且在其中分别判断p != nil和q != nil，再进行对应的加总和指针移动，
是比较简洁的做法。
*/

/**
之前的写法，比较冗余：
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, q, r := l1, l2, res
	var sum, carry int
	for p != nil && q != nil {
		sum = p.Val + q.Val + carry
		carry = sum / 10
		r.Next = &ListNode{Val: sum % 10}
		p, q, r = p.Next, q.Next, r.Next
	}
	for p != nil {
		sum = p.Val + carry
		carry = sum / 10
		r.Next = &ListNode{Val: sum % 10}
		p, r = p.Next, r.Next
	}
	for q != nil {
		sum = q.Val + carry
		carry = sum / 10
		r.Next = &ListNode{Val: sum % 10}
		q, r = q.Next, r.Next
	}
	if carry > 0 { // 还把这段逻辑漏掉了，导致丢失了最后的进位……
		r.Next = &ListNode{Val: carry}
	}
	return res.Next
}
*/
