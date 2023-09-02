package main

import (
	"fmt"

	"github.com/LorenzoSaliu/dataStruct/types"
)

func main() {
	var gm types.Graph_M
	vertex := []*types.Node{
		{Data: 1},
		{Data: 2},
		{Data: 3},
		{Data: 4},
		{Data: 5},
	}
	gm.NewGraphM(vertex)

	fmt.Println(gm)
	gm.AddNode(&types.Node{Data: 6})

	fmt.Println(gm)

}

func populateList(myList *types.LinkedList) {
	var i int
	for i = 0; i < 10; i++ {
		node := types.Node{Data: fmt.Sprint("node-", i)}
		myList.AddFirst(&node)
	}

	for ; i < 20; i++ {
		node := types.Node{Data: fmt.Sprint("node-", i)}
		myList.AddLast(&node)
	}
}

func populateQueue(q *types.Queue) {
	for i := 0; i < 10; i++ {
		node := types.Node{Data: fmt.Sprint("node-", i)}
		q.Push(&node)
	}
}
func populateStack(s *types.Stack) {
	for i := 0; i < 10; i++ {
		node := types.Node{Data: fmt.Sprint("node-", i)}
		s.Push(&node)
	}
}
