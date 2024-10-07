/*
Given the root of a BST, return the minimum absolute difference between the values of any two different nodes in the tree.
*/
package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}
type BinarySearchTree struct {
	Root *Node
}

func main() {
	t := BinarySearchTree{Root: &Node{Val: 11}}
	t.Root.Left = &Node{Val: 5}
	t.Root.Right = &Node{Val: 15}
	t.Root.Left.Left = &Node{Val: 1}
	t.Root.Left.Right = &Node{Val: 9}
	var prev *Node
	diff := math.MaxInt
	fmt.Println(absMin(t.Root, &prev, &diff))

}
func absMin(n *Node, prev **Node, diff *int) int {

	if n == nil {
		return 0
	}

	absMin(n.Left, prev, diff)
	fmt.Println(n, prev, diff)

	if *prev != nil {
		//fmt.Println("node and prev", n.Val, prev.Val)
		d := n.Val - (*prev).Val
		if d < *diff {
			*diff = d
			//	fmt.Println("diff", diff)
		}
	}

	*prev = n

	absMin(n.Right, prev, diff)

	return *diff
}

func min(a, b int) int {
	fmt.Println("numbers for min", a, b)
	if a < b {
		return a
	}
	return b
}
