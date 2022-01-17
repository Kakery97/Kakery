package LinkedList

// ListNode - 单向链表结构体定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// DuListNode - 双向链表结构体定义
type DuListNode struct {
	Val  int
	Prev *DuListNode
	Next *DuListNode
}
