package types

import "fmt"

type Stack struct {
	Nodes  []*Node
	Length int
}

func (s *Stack) Push(n *Node) {
	s.Nodes = append(s.Nodes, n)
	s.Length++
}

func (s *Stack) Pop() *Node {
	if s.Length == 0 {
		return nil
	}

	n := s.Nodes[s.Length-1:]
	s.Nodes = s.Nodes[:s.Length-1]
	s.Length--
	return n[0]
}

func (s Stack) String() string {
	st := "{ "
	for i := 0; i < s.Length; i++ {
		if i != s.Length-1 {
			st += fmt.Sprint(s.Nodes[i]) + ", "
		} else {
			st += fmt.Sprint(s.Nodes[i])
		}
	}
	return st + " }"
}
