package types

import (
	"fmt"
)

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) AddFirst(n *Node) {
	n.Next = l.Head
	l.Head = n
	l.Length++
}

func (l *LinkedList) AddLast(n *Node) {
	next := l.Head
	for ; next.Next != nil; next = next.Next {
	}
	next.Next = n
	l.Length++
}

func (l LinkedList) GetNode(n *Node) *Node {
	for ; l.Head != nil; l.Head = l.Head.Next {
		if l.Head.Data == n.Data {
			return l.Head
		}
	}
	return nil
}

func (l *LinkedList) RemoveNode(n *Node) {
	var previus *Node
	current := l.Head

	// head nil, nothing to remove in empty list
	if current == nil {
		return
	}
	// node to remove is head
	if current.Data == n.Data {
		//list of only head
		if current.Next == nil {
			l.Head = nil
		} else {
			l.Head = current.Next
		}
		l.Length--
		return
	}
	// node to remove is not head
	for ; current != nil; current = current.Next {
		if current.Data == n.Data {
			previus.Next = current.Next
			l.Length--
			return
		}
		previus = current
	}
}

func (l LinkedList) String() string {
	s := "{ "
	for l.Head != nil {
		if l.Head.Next != nil {
			s += fmt.Sprint(l.Head) + ", "
		} else {
			s += fmt.Sprint(l.Head)
		}
		l.Head = l.Head.Next
	}
	return s + " }"
}
