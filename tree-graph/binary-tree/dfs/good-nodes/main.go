//Given the root of a binary tree, find the number of nodes that are good. A node is good if the path between the root and the node has no nodes with a greater value.

/*
			5
			/\
		4		8
		/		/\
	11		  13	4
	/\				\
	7  2			  1

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

type BinaryTree struct {
	Root *Node
}

func main() {
	t := &BinaryTree{Root: &Node{Val: 5}}
	t.Root.Left = &Node{Val: 4}
	t.Root.Left.Left = &Node{Val: 11}
	t.Root.Left.Left.Left = &Node{Val: 7}
	t.Root.Left.Left.Right = &Node{Val: 2}
	t.Root.Right = &Node{Val: 8}
	t.Root.Right.Left = &Node{Val: 13}
	t.Root.Right.Right = &Node{Val: 4}
	t.Root.Right.Right.Right = &Node{Val: 1}

	fmt.Println(calcGoodNode(t.Root, int(math.Inf(-1))))

}

func calcGoodNode(n *Node, maxSoFar int) int {
	if n == nil {
		return 0
	}

	left := calcGoodNode(n.Left, max(n.Val, maxSoFar))
	right := calcGoodNode(n.Right, max(n.Val, maxSoFar))
	ans := left + right
	if n.Val >= maxSoFar {
		ans += 1
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
