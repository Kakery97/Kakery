package BinaryTree

// 二叉树的最近公共祖先 (LeetCode_236)
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 使用二叉树递归实现判断 (树形DP / 后序遍历解法)
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 终止条件:
	// 1. 当越过叶节点，则直接返回 null
	// 2. 当 root 等于 p 或者 q 则直接返回 root
	if root == nil || root == p || root == q {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	// 返回值:
	// 1. 当 left 和 right 同时为空, 说明 root 的左 / 右子树中都不包含 p 或者 q, 返回 null
	// 2. 当 left 和 right 同时不为空, 说明 p, q 分列在 root 的左 / 右子树, 因此 root 为最近公共祖先, 返回 root
	// 3. 当 left 为空, right 不为空, p, q 都不在 root 的左子树中, 直接返回 right, 具体可分为两种情况:
	//    1). p, q 其中一个在 root 的右子树中, 此时 right 指向 p 或者 q
	//    2). p, q 两节点都在 root 的右子树中, 此时的 right 指向最近公共祖先节点
	// 4. 当 left 不为空, right 为空, 与情况 3. 同理
	if left != nil && right != nil {
		return root
	}
	// 合并情况 1, 3, 4
	if left != nil {
		return left
	} else {
		return right
	}
}
