package Graph

// 图的宽度优先遍历
func bfs(node *Node) (res []int) {
	if node == nil {
		return
	}
	queue := []*Node{node}
	set := map[*Node]bool{node: true}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur.Val)
		for _, n := range cur.Next {
			if !set[n] {
				set[n] = true
				queue = append(queue, n)
			}
		}
	}
	return
}
