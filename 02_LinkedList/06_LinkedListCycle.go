package LinkedList

// 环形链表判断 (LeetCode_141, LeetCode_142)
// https://leetcode.cn/problems/linked-list-cycle/  -  判断链表中是否有环
// https://leetcode.cn/problems/linked-list-cycle-ii/  -  返回链表开始入环的第一个节点
// 实现方法: 采用快慢指针实现链表环形的判断
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow { // 慢指针追上快指针, 说明有环
			return true
		}
	}
	return false // 链表正常遍历完毕, 说明无环
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			fast = head        // 1. 慢指针追上快指针后, 将快指针放到头部位置, 慢指针位置保持不变
			for fast != slow { // 2. 快指针同时也每次走一步, 快慢指针同时遍历
				slow = slow.Next
				fast = fast.Next
			}
			return fast // 3. 快慢指针第二次相遇的位置就是入环的第一个节点
		}
	}
	return nil
}
