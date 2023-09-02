package types

import "fmt"

type Node struct {
	Data any
	Next *Node
}

func (n Node) String() string {
	return fmt.Sprint(n.Data)
}
