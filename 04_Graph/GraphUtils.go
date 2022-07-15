package Graph

// Node - 代表图的结点的结构体
type Node struct {
	Val     int
	In, Out int
	Next    []*Node
	Edges   []*Edge
}

// Edge - 代表图的边的结构体
type Edge struct {
	Weight   int
	From, To *Node
}

// Graph - 图的结构体定义
type Graph struct {
	Nodes map[int]*Node
	Edges map[*Edge]bool
}
