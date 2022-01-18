package LinkedList

// 相交链表起始节点判断 (LeetCode_160)
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// 找出并返回两个单链表相交的起始节点
// 实现方法: 采用快慢指针实现链表环形的判断
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
// LeetCode原题 - 保证两个链表中都不存在环的情况
func getIntersectionNodeNoLoop(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	endA, endB := headA, headB
	for endA.Next != nil {
		lenA++
		endA = endA.Next
	}
	for endB.Next != nil {
		lenB++
		endB = endB.Next
	}
	if endA != endB { // 1. 若两个无环链表的最后一个节点不同, 必定不相交
		return nil
	}
	if lenA > lenB { // 2. 计算差值使两个链表长度相等
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}
	for headA != headB { // 3. 同时遍历, 寻找相交点(前面已经保证相交点一定存在)
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

// 附 - 两个链表中都存在环的情况
// loopNodeA, loopNodeB分别为两个链表的第一个入环节点
func getIntersectionNodeBothLoop(headA, headB, loopNodeA, loopNodeB *ListNode) *ListNode {
	curA := loopNodeA.Next      // 零时节点设置为链表A第一个入环节点的下一个
	if loopNodeA == loopNodeB { // 入环节点相同, 说明一定相交且共用入环节点
		loopNodeA.Next = nil // 去掉环就等价于两个链表中都不存在环的情况
		res := getIntersectionNodeNoLoop(headA, headB)
		loopNodeA.Next = curA // 还原整体链表结构
		return res
	} else { // 入环节点不同, 链表可能相交也可能不相交
		for curA != loopNodeA { // 环A遍历一圈, 若是能碰到链表B第一个入环节点, 就说明有相交点
			if curA == loopNodeB {
				return loopNodeA
			}
			curA = curA.Next
		}
		return nil
	}
}

// 附 - 总情况: 给定两个可能有环也可能无环的单链表, 找出并返回两个单链表相交的起始节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	loopNodeA, loopNodeB := detectCycle(headA), detectCycle(headB)
	if loopNodeA == nil && loopNodeB == nil { // 1. 两个链表中都不存在环的情况
		return getIntersectionNodeNoLoop(headA, headB)
	} else if loopNodeA == nil && loopNodeB == nil { // 2. 两个链表中都存在环的情况
		return getIntersectionNodeBothLoop(headA, headB, loopNodeA, loopNodeB)
	} else { // 3. 只有其中一个链表有环, 一定不相交
		return nil
	}
}
