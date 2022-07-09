package BinaryTree

// 深度优先遍历(Depth-First-Search, DFS), 就是从一个端点一直查找到另一个端点, "一头深入"
// 在二叉树的遍历中, 深度优先遍历分为前序遍历、中序遍历以及后序遍历

// 1. 二叉树的前序遍历 (LeetCode_114)
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
// 分别使用递归以及非递归实现
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
// 递归1 - 按照遍历的思维
func preorderTraversal1(root *TreeNode) (res []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 遍历法一般使用全局变量存储结果
		// 前序：刚刚进入当前节点
		res = append(res, node.Val)
		// 遍历左子树
		dfs(node.Left)
		// 遍历右子树
		dfs(node.Right)
	}
	// 遍历根节点
	dfs(root)
	return
}

// 递归2 - 按照分解子问题的思维
func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	// 前序：刚刚进入当前节点
	res = append(res, root.Val)
	// 加入左子树处理完的结果
	res = append(res, preorderTraversal(root.Left)...)
	// 加入右子树处理完的结果
	res = append(res, preorderTraversal(root.Right)...)
	return
}

// 迭代1 - 按照头->右->左入栈顺序的思维
func preorderTraversalNonRecursive1(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val) // 出栈时打印
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}

// 迭代2 - 按照头/左树切割的思维
func preorderTraversalNonRecursive2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val) // 在入栈时打印即为前序遍历
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return
}

// 2. 二叉树的中序遍历 (LeetCode_94)
// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 分别使用递归以及非递归实现
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
// 递归1 - 按照遍历的思维
func inorderTraversal1(root *TreeNode) (res []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

// 递归2 - 按照分解子问题的思维
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return
}

// 迭代 - 按照头/左树切割的思维
func inorderTraversalNonRecursive(root *TreeNode) (res []int) {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1] // 在出栈时打印即为中序遍历
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
	}
	return
}

// 3. 二叉树的后序遍历 (LeetCode_145)
// https://leetcode.cn/problems/binary-tree-postorder-traversal/
// 分别使用递归以及非递归实现
//
// 时间复杂度  О(n)
// 空间复杂度  O(n)
// 递归1 - 按照遍历的思维
func postorderTraversal1(root *TreeNode) (res []int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return
}

// 递归2 - 按照分解子问题的思维
func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return
}

// 迭代1 - 使用两个栈按, 照头->左->右入栈顺序, 在依次入第二个栈的思维
func postorderTraversalNonRecursive1(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	stack1 := []*TreeNode{root}
	var stack2 []int
	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node.Val)
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}
	for len(stack2) > 0 { // 出第二个栈时打印
		num := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		res = append(res, num)
	}
	return
}

// 迭代2 - 按照头/左树切割的思维
// 二叉树的公共非递归方法难点其实就在于后序遍历
// 由于根结点需要最后访问, 于是根节点不能先弹出不管, 所以后序需要两次路过根节点, 因此遍历的时候需要根据第几次路过来决定是否访问根节点
// 解决的办法是增加一个prev指针, 指向上次访问的节点
// 因为后序遍历最后一定是遍历根节点, 所以当prev指向root.Right的时候那么就是该遍历root的时候了
func postorderTraversalNonRecursive2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	var prev *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if node.Right != nil && prev != node.Right { // 该节点的右孩子不空, 而且上一个访问的不是右孩子(证明这是从左孩子回溯过来的)
			node = node.Right
		} else { // 该节点的右孩子为空, 或者右孩子已经访问过了
			res = append(res, node.Val) // 遍历节点
			stack = stack[:len(stack)-1]
			prev = node
			node = nil // 防止node再次被压入堆栈, 所以要置空
		}
	}
	return
}
