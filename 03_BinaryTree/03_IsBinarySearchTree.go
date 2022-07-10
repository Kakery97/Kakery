package BinaryTree

import "math"

// 二叉搜索树(Binary Search Tree, BST), 是指一棵空树或者具有下列性质的二叉树:
// 1. 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值
// 2. 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值
// 3. 任意节点左、右子树也分别为二叉搜索树

// 验证二叉搜索树 (LeetCode_98)
// https://leetcode.cn/problems/validate-binary-search-tree/
// 使用中序遍历是否递增实现判断
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
func isValidBST1(root *TreeNode) bool {
	prev := math.MinInt64
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		isLeftValidBST := dfs(node.Left)
		if !isLeftValidBST {
			return false
		}
		if prev < node.Val {
			prev = node.Val
		} else {
			return false
		}
		return dfs(node.Right)
	}
	return dfs(root)
}

// 附 - 使用根据定义的二叉树递归实现
func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
