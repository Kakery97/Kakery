package LinkedList

// 反转单向链表 (LeetCode_206)
// https://leetcode-cn.com/problems/reverse-linked-list/
// 采用双指针迭代实现
// 因为链表头会改变所以会有返回值
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		prev = head
		head = head.Next
		prev.Next = next
		next = prev
	}

	return prev
}

// 附 - 反转双向链表
func reverseDuList(head *DuListNode) *DuListNode {
	var prev, next *DuListNode
	for head != nil {
		prev = head
		head = head.Next
		prev.Next = next
		prev.Prev = head
		next = prev
	}

	return prev
}
