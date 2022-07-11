package BinaryTree

import (
	"strconv"
	"strings"
)

// 二叉树的序列化和反序列化 (LeetCode_297)
// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
// 比较方便的解法: 利用二叉树的前序遍历以及层序遍历
// 中序序列化的结果和二叉树的结果不是一一对应的, 因此不存在中序遍历的解法
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
// 1. 使用前序遍历, 因为根->左->右的打印顺序在反序列化时更容易定位出根节点的值
func serialize1(root *TreeNode) (res string) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res += "#_"
			return
		}
		res += strconv.Itoa(node.Val) + "_"
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return
}

// 反序列化需要先按照前序遍历找到头然后分别dfs递归得到左右结点
// 由于取结点是需要先进先出的顺序, 因此采用队列存储
func deserialize1(data string) *TreeNode {
	// 由于后序遍历时顺序为左->右->根, 所以需要换成栈来实现, 同时恢复的顺序为头->右->左
	nodeQueue := strings.Split(data, "_")
	var dfs func(nodeList *[]string) *TreeNode
	dfs = func(nodeList *[]string) *TreeNode {
		nodeStr := (*nodeList)[0]
		*nodeList = (*nodeList)[1:]
		if nodeStr == "#" {
			return nil
		}
		val, _ := strconv.Atoi(nodeStr)
		return &TreeNode{
			Val:   val,
			Left:  dfs(nodeList),
			Right: dfs(nodeList),
		}
	}
	return dfs(&nodeQueue)
}

// 2. 使用层序遍历
func serialize2(root *TreeNode) (res string) {
	if root == nil {
		return "#_"
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res += "#_"
			continue
		}
		res += strconv.Itoa(node.Val) + "_"
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return
}

func deserialize2(data string) *TreeNode {
	nodeQueue := strings.Split(data, "_")
	res := generateNodeByString(nodeQueue[0])
	if res == nil {
		return nil
	}
	nodeQueue = nodeQueue[1:]
	queue := []*TreeNode{res}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		left, right := generateNodeByString(nodeQueue[0]), generateNodeByString(nodeQueue[1])
		nodeQueue = nodeQueue[2:]
		if left != nil {
			node.Left = left
			queue = append(queue, node.Left)
		}
		if right != nil {
			node.Right = right
			queue = append(queue, node.Right)
		}
	}
	return res
}

func generateNodeByString(str string) *TreeNode {
	if str == "#" {
		return nil
	}
	val, _ := strconv.Atoi(str)
	return &TreeNode{Val: val}
}
