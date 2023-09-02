package types

type Graph_M struct {
	NodeNumber int
	Nodes      []*Node
	Edge       [][]bool
}

func (gm *Graph_M) NewGraphM(n []*Node) {
	gm.NodeNumber = len(n)
	gm.Nodes = n
	gm.Edge = make([][]bool, len(n))
	for i := range gm.Edge {
		gm.Edge[i] = make([]bool, len(n))
	}
}

func (gm *Graph_M) AddNode(n *Node) {
	gm.NodeNumber++
	gm.Nodes = append(gm.Nodes, n)
	//edge := gm.Edge
	gm.Edge = make([][]bool, gm.NodeNumber)
	for i := range gm.Edge {
		gm.Edge[i] = make([]bool, gm.NodeNumber)
	}
}
func (gm *Graph_M) RemoveNode(n *Node) {}
func (gm *Graph_M) AddEdge()           {}
func (gm *Graph_M) RemoveEdge()        {}
func (gm *Graph_M) GetNode(n *Node)    {}
