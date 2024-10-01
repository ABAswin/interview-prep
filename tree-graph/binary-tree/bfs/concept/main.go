package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}
type BinaryTree struct {
	Root *Node
}

type Queue []*Node

func main() {
	t := &BinaryTree{Root: &Node{Val: 1}}
	t.Root.Left = &Node{Val: 2}
	t.Root.Right = &Node{Val: 3}
	t.Root.Left.Left = &Node{Val: 4}
	t.Root.Left.Right = &Node{Val: 5}
	t.Root.Right.Left = &Node{Val: 6}
	t.Root.Right.Right = &Node{Val: 7}

	var q Queue
	q = append(q, t.Root)
	listTree(q)

}
func listTree(q Queue) {
	l := 0
	for len(q) != 0 {

		for _, v := range q {
			q = q[1:]
			fmt.Println("level", l, "val", v.Val)

			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}

		}
		l++
	}
}
