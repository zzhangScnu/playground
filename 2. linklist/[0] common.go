package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

type TwoWayListNode struct {
	Val  int
	Next *TwoWayListNode
	Pre  *TwoWayListNode
}
