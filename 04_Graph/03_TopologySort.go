package Graph

// 拓扑排序(Topological sort)
//
// 有向图的拓扑排序是对其顶点的一种线性排序, 使得对于从顶点 u 到顶点 v 的每个有向边 u->v , u 在排序中都在 v 之前
// 算法要求: 1-有向图; 2=有入度为0的结点; 3-图没有环
func topologicalSort(graph *Graph) (res []int) {
	if graph == nil {
		return
	}
	inMap := map[*Node]int{} // value为某个结点剩余的入度
	var zeroQueue []*Node    // 入度为0的结点进入此队列
	for _, node := range graph.Nodes {
		inMap[node] = node.In
		if node.In == 0 {
			zeroQueue = append(zeroQueue, node)
		}
	}
	for len(zeroQueue) > 0 {
		cur := zeroQueue[0]
		zeroQueue = zeroQueue[1:]
		res = append(res, cur.Val)
		for _, n := range cur.Next {
			inMap[n] = inMap[n] - 1
			if inMap[n] == 0 {
				zeroQueue = append(zeroQueue, n)
			}
		}
	}
	return
}
