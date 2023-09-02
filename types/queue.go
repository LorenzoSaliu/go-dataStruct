package types

import "fmt"

type Queue struct {
	Nodes  []*Node
	Length int
}

func (q *Queue) Push(n *Node) {
	q.Nodes = append(q.Nodes, n)
	q.Length++
}

func (q *Queue) Pop() *Node {
	if q.Length == 0 {
		return nil
	}

	n := q.Nodes[:1]
	q.Nodes = q.Nodes[1:]
	q.Length--
	return n[0]
}

func (q Queue) String() string {
	s := "{ "
	for i := 0; i < q.Length; i++ {
		if i != q.Length-1 {
			s += fmt.Sprint(q.Nodes[i]) + ", "
		} else {
			s += fmt.Sprint(q.Nodes[i])
		}
	}
	return s + " }"
}
