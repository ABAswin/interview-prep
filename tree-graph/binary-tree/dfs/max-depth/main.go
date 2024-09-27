// find the max depth of binary tree
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

/*
			2
			/\
		5		3
		/\
	7		8
			\
			9
*/
func main() {
	t := &BinaryTree{}
	t.Root = &Node{Val: 2}
	t.Root.Left = &Node{Val: 5}
	t.Root.Right = &Node{Val: 3}
	t.Root.Left.Left = &Node{Val: 7}
	t.Root.Left.Right = &Node{Val: 8}
	t.Root.Left.Right.Right = &Node{Val: 9}

	fmt.Println(CalculateDepth(t.Root))
}

func CalculateDepth(n *Node) int {
	// base case return 0
	if n == nil {
		return 0
	}

	left := CalculateDepth(n.Left)
	right := CalculateDepth(n.Right)

	return 1 + max(left, right)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
