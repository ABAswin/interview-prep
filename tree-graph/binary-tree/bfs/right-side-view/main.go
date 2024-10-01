/*
Given the root of a binary tree, imagine yourself standing on the right side of it. Return the values of the nodes you can see ordered from top to bottom.
*/

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
	t.Root.Right.Left = &Node{Val: 4}
	t.Root.Right.Right = &Node{Val: 5}
	// t.Root.Right.Left = &Node{Val: 6}
	// t.Root.Right.Right = &Node{Val: 7}

	var q Queue
	q = append(q, t.Root)
	fmt.Println(rightMostNodes(q))
}

func rightMostNodes(q Queue) []int {
	var ret []int
	for len(q) != 0 {
		ret = append(ret, q[len(q)-1].Val)
		for _, v := range q {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			q = q[1:]
		}

	}
	return ret
}
