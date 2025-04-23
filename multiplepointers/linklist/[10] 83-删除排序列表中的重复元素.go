package linklist

import "code.byted.org/zhanglihua.river/playground/linklist"

func deleteDuplicates(head *linklist.ListNode) {
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
}

/**
思路：
跟数组版的删除重复元素基本一致。
唯一要注意的是，最后需要slow.Next = nil，即将调整后的子链表的末尾指向nil，
防止指向重复的元素。
*/
