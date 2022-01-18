package LinkedList

// 复制带随机指针的链表 (LeetCode_138)
// https://leetcode-cn.com/problems/copy-list-with-random-pointer/
// 链表每个节点包含一个额外增加的随机指针random, 该指针可以指向链表中的任何节点或空节点
// 构造这个链表的深拷贝
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
// 实现方法: 迭代 + 节点拆分
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil { // 1. 将该链表中每一个节点拆分为两个相连的节点
		temp := &Node{ // 新建结点, 不复制前一个结点的随机指针random
			Val:    cur.Val,
			Next:   cur.Next,
			Random: nil,
		}
		cur.Next = temp
		cur = cur.Next.Next
	}
	cur = head
	for cur != nil { // 2. 遍历链表, 可以直接找到每一个拷贝节点的随机指针应当指向的节点random->Next
		if cur.Random != nil { // *注意原节点的随机指针可能为空!
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	newHead := head.Next
	cur = head
	for cur != nil { // 3. 将这个链表按照原节点与拷贝节点的种类进行拆分
		temp := cur.Next
		cur.Next = cur.Next.Next
		if temp.Next != nil { // *注意最后一个拷贝节点的后节点可能为空!
			temp.Next = temp.Next.Next
		}
		cur = cur.Next
	}
	return newHead
}
