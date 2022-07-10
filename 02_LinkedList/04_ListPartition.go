package LinkedList

// 分隔链表 (LeetCode_86)
// https://leetcode.cn/problems/partition-list/
// 使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前
// 同时保留两个分区中每个节点的初始相对位置 (保持稳定)
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
// 1. 不使用哑节点的情况, 需要自行判断边界条件
func partition1(head *ListNode, x int) *ListNode {
	var leftHead, leftTail, rightHead, rightTail *ListNode
	for head != nil {
		if head.Val < x {
			if leftHead == nil { // 遇到第一个先初始化头尾结点
				leftHead = head
				leftTail = head
			} else {
				leftTail.Next = head
				leftTail = leftTail.Next
			}
		} else {
			if rightHead == nil {
				rightHead = head
				rightTail = head
			} else {
				rightTail.Next = head
				rightTail = rightTail.Next
			}
		}
		head = head.Next
	}
	if leftHead != nil && rightHead != nil { // 边界情况判断
		rightTail.Next = nil
		leftTail.Next = rightHead
		return leftHead
	} else if leftHead != nil && rightHead == nil {
		return leftHead
	} else if leftHead == nil && rightHead != nil {
		return rightHead
	} else {
		return nil
	}
}

// 2. 使用哑节点的情况, 可以更方便地处理头节点为空的边界条件, 代码更加简洁
func partition2(head *ListNode, x int) *ListNode {
	leftHead, rightHead := &ListNode{}, &ListNode{} // 空结点作为两边的头结点
	leftTail, rightTail := leftHead, rightHead
	for head != nil {
		if head.Val < x {
			leftTail.Next = head
			leftTail = leftTail.Next
		} else {
			rightTail.Next = head
			rightTail = rightTail.Next
		}
		head = head.Next
	}
	rightTail.Next = nil
	leftTail.Next = rightHead.Next
	return leftHead.Next
}

// 附 - partition划分为左小于, 中间等于, 右边大于三个区域
func partition(head *ListNode, x int) *ListNode {
	leftHead, midHead, rightHead := &ListNode{}, &ListNode{}, &ListNode{}
	leftTail, midTail, rightTail := leftHead, midHead, rightHead
	for head != nil {
		if head.Val < x {
			leftTail.Next = head
			leftTail = leftTail.Next
		} else if head.Val == x {
			midTail.Next = head
			midTail = midTail.Next
		} else {
			rightTail.Next = head
			rightTail = rightTail.Next
		}
		head = head.Next
	}
	rightTail.Next = nil
	leftTail.Next = midHead.Next
	midTail.Next = rightHead.Next
	return leftHead.Next
}
