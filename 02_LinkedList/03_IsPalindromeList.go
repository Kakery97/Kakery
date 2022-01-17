package LinkedList

// 判断链表是否为回文链表
// 采用快慢指针实现
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var mid, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil { // 1.通过快慢指针找到链表中点
		mid = mid.Next
		fast = fast.Next.Next
	}
	var cur, perv *ListNode = mid.Next, nil
	fast = nil
	mid.Next = nil
	for cur != nil { // 2.将后半部分的链表反转, fast作为右链表的头
		fast = cur
		cur = cur.Next
		fast.Next = perv
		perv = fast
	}
	cur, perv = head, fast
	for cur != nil && perv != nil { // 3.判依次遍历断是否是回文链表
		if cur.Val != perv.Val {
			return false
		} else {
			cur = cur.Next
			perv = perv.Next
		}
	}
	cur, perv = nil, nil
	for fast != nil { // 4.还原链表
		cur = fast
		fast = fast.Next
		cur.Next = perv
		perv = cur
	}
	mid.Next = cur
	return true
}
