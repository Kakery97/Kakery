package BinaryTree

// TreeNode - 二叉树节点结构体定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node - 带有父结点引用的二叉树节点结构体定义
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
