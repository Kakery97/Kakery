package BinaryTree

// 平衡二叉树是一种每个节点的左右两子树高度差都不超过一的二叉树

// 平衡二叉树 (LeetCode_110)
// https://leetcode.cn/problems/balanced-binary-tree/
// 使用二叉树递归实现判断 (树形DP解法)
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
func isBalanced(root *TreeNode) bool {
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := helper(node.Left), helper(node.Right)
		// 负数表示非平衡二叉树: 左或右子树非平衡二叉树或者子树高度差超过了1
		if l < 0 || r < 0 || l-r > 1 || l-r < -1 {
			return -1
		}
		return maxInt(l, r) + 1
	}
	return helper(root) >= 0
}
