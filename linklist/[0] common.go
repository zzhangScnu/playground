package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

type TwoWayListNode struct {
	Key  int
	Val  int
	Next *TwoWayListNode
	Pre  *TwoWayListNode
}
