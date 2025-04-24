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

var left *ListNode

func isPalindromeRecursively(head *ListNode) bool {
	left = head
	return doIsPalindromeRecursively(head)
}

func doIsPalindromeRecursively(right *ListNode) bool {
	if right == nil {
		return true
	}
	isEqual := doIsPalindromeRecursively(right.Next)
	isEqual = isEqual && (left.Val == right.Val)
	left = left.Next
	return isEqual
}

/**
暴力解法1：
遍历链表 -> 转换为数组 -> 双指针判断是否回文数组
时间复杂度：O(n)
空间复杂度：O(n)


暴力解法2：
反转链表 -> 比较正反链表是否相等
时间复杂度：O(n)
空间复杂度：O(n)


递归解法：isPalindromeRecursively
类似二叉树的遍历，可以将链表看作无分叉的、每个节点只有一个子节点的树，不同的遍历顺序代表不同的输出顺序。
如：
func traverse(node *ListNode) {
	if node == nil {
		return node
	}
	traverse(node.Next)
	fmt.Println(node.Val)
}
后序遍历是在递归下钻，直至叶子节点的下一层后触底反弹，在反弹过程中输出节点值。
核心逻辑是将链表节点依次放入栈，再依次弹出，所以输出顺序是逆序的。
同理，如果是先序遍历，则为正序，跟迭代遍历链表效果一致。
那么可以利用后序遍历特性，在递归的逆序访问中，将当前节点right，与正序访问的节点left，进行比较，如果不一致，则不是回文链表。
注意，递归结果需要变量接收，且与当前层的比较结果做与运算。本质上是left和right从两边向中间夹逼，直至left遍历到链表终点，right遍历到链表起点。
时间复杂度：O(n)
空间复杂度：O(n)，需要额外堆栈支持。


原地解法：isPalindrome
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
