/*Given the root of a binary tree, find the maximum value v for which there exist different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.

A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.*/

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
	t := BinaryTree{Root: &Node{Val: 1}}
	t.Root.Right = &Node{Val: 2}
	t.Root.Right.Right = &Node{Val: 0}
	t.Root.Right.Right.Left = &Node{Val: 3}
	// t.Root.Right.Left = &Node{Val: 3}
	// if the above 3 was a child node of 2 along with 0, the output will be different as the difference is always calculated with ancestor nodes
	// Learning: in an edge case where root node is nil, simply return zero
	// here we are passing non nil tree
	fmt.Println(maxDiffAnc(t.Root, math.MaxInt32, math.MinInt32))

}
func maxDiffAnc(n *Node, minVal, maxVal int) int {
	if n == nil {
		return maxVal - minVal
	}

	minVal = min(n.Val, minVal)
	maxVal = max(n.Val, maxVal)
	fmt.Println("n val", n.Val)
	fmt.Println("The min and max are", minVal, maxVal)

	// maxdif so far
	// min of node val

	left := maxDiffAnc(n.Left, minVal, maxVal)
	right := maxDiffAnc(n.Right, minVal, maxVal)

	return max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
