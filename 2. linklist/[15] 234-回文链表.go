package linklist

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：head = [1,2,2,1]
// 输出：true
//
// 示例 2：
//
// 输入：head = [1,2]
// 输出：false
//
// 提示：
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	p, q := head, reverse(slow)
	for q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/**
暴力解法：
遍历链表 -> 转换为数组 -> 双指针判断是否回文数组
时间复杂度：O(n)
空间复杂度：O(n)

原地解法：
遍历链表 -> 找到中间节点 -> 反转后半段链表 -> 从头开始对比两段链表是否相等
时间复杂度：O(n)
空间复杂度：O(1)
需要注意的细节：
1. 反转链表时，用2个指针操作即可，分别指向前一个节点和当前节点。
   因反转结束的条件是cur == nil，此时pre指向的是原链表中的最后一个节点，即反转后当前的头节点；
2. 反转后，子链表p的范围为[head, slow]，q的范围为[tail, slow]。p和q有2种情况：
	- 原链表节点数量为奇数：1->2->3->2->1，p = 1->2->3，q = 1->2->3；
	- 原链表节点数量为偶数：1->2->2->1，p = 1->2，q = 1->2。
   即len(p) == len(q)。但因为p的末尾仍连接在原链表的下一个节点上，而没有显式处理指向nil，
   所以应遍历q，同时推进p和q的指针。
*/
