package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) insert(val int) {
	// var temp Node
	// temp.Data = val
	temp := &Node{Data: val}
	//fmt.Printf("address of temp %p\n", temp)
	if l.Head == nil {
		fmt.Print("Setting head\n")
		l.Head = temp
		//fmt.Printf("%p", *l.Head)
		//l.print()
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = temp
}

func (l *List) print() {
	fmt.Print("printing list\n")
	i := l.Head
	for i != nil {
		fmt.Printf("%d\n", i.Data)
		i = i.Next
	}
}

func (l *List) delete(val int) {
	// 1,2,3
	// 1
	// 1 2
	if l.Head.Data == val {
		l.Head = l.Head.Next
		return
	}
	i := l.Head
	for i != nil {
		if i.Data == val {
			l.Head = i.Next
			i = i.Next
			continue
		}

		if i.Next.Data == val {
			i.Next = i.Next.Next
		}
	}

}

func main() {

	l := &List{}
	l.insert(1)
	fmt.Printf("after inserting 1 %p\n", l.Head)
	l.print()
	l.insert(2)
	l.print()
	l.insert(2)
	l.insert(2)
	l.insert(2)
	l.insert(2)

	l.print()
}
