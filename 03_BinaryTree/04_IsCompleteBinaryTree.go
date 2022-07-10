package BinaryTree

// 完全二叉树(Complete Binary Tree, CBT), 是指一棵空树或者具有下列性质的二叉树:
// 在二叉树中, 若除最后一层外的其余层都是满的, 并且最后一层要么是满的, 要么在右边缺少若干连续节点, 则此二叉树为完全二叉树

// 二叉树的完全性检验 (LeetCode_958)
// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
// 使用层序遍历实现判断
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
func isCompleteTree1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isLeafCheck := false // 是否开启判断叶子结点
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		l, r := node.Left, node.Right
		// 不是完全二叉树的两种情况:
		// 1. 该是叶子结点的时候有了自结点
		// 2. 某结点出现左孩子为空但是有右孩子的情况
		if (isLeafCheck && (l != nil || r != nil)) || (l == nil && r != nil) {
			return false
		}
		if l != nil {
			queue = append(queue, l)
		}
		if r != nil {
			queue = append(queue, r)
		} else {
			isLeafCheck = true
		}
	}
	return true
}

// 附 - 队列中传空的做法
// 什么时候是不完全呢? 其实就是出现空结点之后后面又出现了结点
func isCompleteTree2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isNilCheck := false // 是否开启判断空结点
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			isNilCheck = true
		} else {
			if isNilCheck && node != nil {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}
