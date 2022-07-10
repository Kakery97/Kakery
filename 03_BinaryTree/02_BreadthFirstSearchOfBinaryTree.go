package BinaryTree

// 广度优先遍历(Breadth-First-Search, BFS), 就是从一个点, 优先向其周围所有的点进行搜索, 类似走迷宫
// 在二叉树中, BFS 实质就是层序遍历

// 二叉树的层序遍历 (LeetCode_102)
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 层序遍历使用队列来辅助实现
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			temp[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return
}

// 附 - 运用题目 二叉树的最大宽度 (LeetCode_662)
// https://leetcode.cn/problems/maximum-width-of-binary-tree/
// 这道题需要统计的最大宽度需要统计空结点
// 使用数组存储二叉树的规律：下标为i的结点, 左孩子下标为2*i+1, 右孩子下标为2*i+2
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
func widthOfBinaryTree(root *TreeNode) (max int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	root.Val = 0
	for len(queue) > 0 {
		temp, size := 0, len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				temp = node.Val
			}
			if i == size-1 {
				max = maxInt(max, node.Val-temp+1)
			}
			if node.Left != nil {
				node.Left.Val = 2*node.Val + 1
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 2
				queue = append(queue, node.Right)
			}
		}
	}
	return
}
