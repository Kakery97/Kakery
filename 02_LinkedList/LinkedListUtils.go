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

// Node - 带有随机结点的链表结构体定义
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
