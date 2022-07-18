package Graph

// 图的深度优先遍历
func dfs(node *Node) (res []int) {
	if node == nil {
		return
	}
	stack := []*Node{node}
	set := map[*Node]bool{node: true}
	res = append(res, node.Val)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, n := range cur.Next {
			if !set[n] {
				set[n] = true
				stack = append(stack, cur)
				stack = append(stack, n)
				res = append(res, n.Val)
				break
			}
		}
	}
	return
}
