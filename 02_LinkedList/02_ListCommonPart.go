package LinkedList

// 返回两个有序链表的公共部分
// 两个链表遍历结束之前: 谁小谁先往后移; 一样的话记录结果并同时往后移
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
func listCommonPart(head1, head2 *ListNode) *ListNode {
	var res, cur *ListNode
	cur = res
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			head1 = head1.Next
		} else if head1.Val == head2.Val {
			cur.Next = &ListNode{
				Val:  head1.Val,
				Next: nil,
			}
			cur = cur.Next
			head1 = head1.Next
			head2 = head2.Next
		} else {
			head2 = head2.Next
		}
	}
	return res.Next
}
