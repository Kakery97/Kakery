package Problems

// LeetCode 366. 寻找二叉树的叶子节点
// https://leetcode.cn/problems/find-leaves-of-binary-tree/
// 二叉树后序遍历的运用 - 反向定义深度
func findLeaves(root *TreeNode) (res [][]int) {
	var dfs func(node *TreeNode) int
	// dfs 采用后序遍历(因为满足自底向上的规律)递归求得当前结点的深度, 把叶子节点作为计算高度的起点, 即叶子结点高度为0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		l, r := dfs(node.Left), dfs(node.Right)
		h := maxInt(l, r) + 1
		// 某层的结点第一次被处理
		if len(res) <= h {
			res = append(res, []int{node.Val})
		} else {
			res[h] = append(res[h], node.Val)
		}
		return h
	}
	dfs(root)
	return
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
