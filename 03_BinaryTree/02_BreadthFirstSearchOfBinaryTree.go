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
