package linklist

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	return reverseListIteratively(head)
	// return reverseListRecursively(nil, head)
	// return reverseListRecursivelyII(head)
}

func reverseListIteratively(head *ListNode) *ListNode {
	cur := head
	var pre, next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseListRecursively(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverseListRecursively(cur, next)
}

func reverseListRecursivelyII(cur *ListNode) *ListNode {
	if cur == nil || cur.Next == nil {
		return cur
	}
	reversed := reverseListRecursivelyII(cur.Next)
	cur.Next.Next = cur
	cur.Next = nil
	return reversed
}

/**
迭代方式：
1. 双指针法，用cur定位当前处理的节点。对于三指针的情况，用pre + cur + next的组合可以省一些边界条件，
   而不是像一开始的想法一样，用cur + curNext + next；
2. next用来保存下一次要处理的节点，否则cur的指针指向一变，就断链了；
3. 当cur为nil时，表示已到达链表末尾，此时返回的pre就是最新的头节点；

递归方式：
【对于递归算法，最重要的就是明确递归函数的定义。】
由此可以有不同的实现方式。
- reverseListRecursively
  1. 接收要处理的两个节点，在每层进行反转操作；
  2. 完成这两个节点的操作后，将后一轮需要反转的节点传递下去；
  3. 当执行base case即cur == nil时，表示已到达链表末尾，此时返回的pre就是最新的头节点。
	  会将这个pre层层向上返回，最终返回给主方法；
- reverseListRecursivelyII
  1. 接收当前节点，并将下一个节点传递到下一次递归方法调用，获取反转后的子链表；
  2. 处理当前节点的指针指向，再返回反转后的头节点。
*/
