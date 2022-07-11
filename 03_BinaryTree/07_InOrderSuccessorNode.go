package BinaryTree

// 二叉搜索树中的中序后继 (LeetCode_285 / 剑指 Offer II 053)
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// https://leetcode.cn/problems/P5rCT8/
// 利用二叉搜索树特性的非中序遍历的解法
//
// 时间复杂度  О(n)
// 空间复杂度  O(1)
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	// 如果节点p的右子树不为空, 则节点p的中序后继在其右子树中, 在其右子树中定位到最左边的节点, 即为节点p的中序后继
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	// 如果节点p的右子树为空, 则需要从根节点开始遍历寻找节点p的祖先节点
	var res *TreeNode
	// 1. 如果node的节点值大于p的节点值, 则p的中序后继可能是node或者在node的左子树中, 因此用node更新答案, 并将node移动到其左子节点继续遍历
	// 2. 如果node的节点值小于或等于p 的节点值, 则p中序后继可能在node的右子树中, 因此将node移动到其右子节点继续遍历
	for root != nil {
		if root.Val > p.Val {
			// 只有在大于时更新res, 可以保证在有中序后继时一定可以找到中序后继, 没有中序后继时res一定为空
			res = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return res
}

// 二叉搜索树中的中序后继 II (LeetCode_510)
// https://leetcode.cn/problems/inorder-successor-in-bst-ii/
// 1. 无法直接访问树
// 2. 每个节点都会有其父节点的引用
//
// 时间复杂度  О(H) - 其中H为数的高度
// 空间复杂度  O(1)
func inorderSuccessor2(node *Node) (res *Node) {
	// 同1, 如果node的右子树不为空, 则node的中序后继在其右子树中, 在其右子树中定位到最左边的节点, 即为node的中序后继
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	// 如果节点p的右子树为空, 则可以根据父结点的引用向上回溯中序后继
	// 找到第一个父结点, 特点为由其左孩子回溯而来, 若找不到则说明没有后继, 返回空
	for node.Parent != nil && node != node.Parent.Left {
		node = node.Parent
	}
	return node.Parent
}
