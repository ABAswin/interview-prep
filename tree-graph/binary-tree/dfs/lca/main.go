//Given the root of a binary tree and two nodes p and q that are in the tree, return the lowest common ancestor (LCA) of the two nodes. The LCA is the lowest node in the tree that has both p and q as descendants (note: a node is a descendant of itself).

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
				3
				/\
			   5   	1
			  / \  	/\
			6	 2  0  8
				/\
				7 4


6,0
7,4
5,1

*/

func main() {
	t := &BinaryTree{}
	t.Root = &Node{Val: 3}
	t.Root.Left = &Node{Val: 5}
	t.Root.Right = &Node{Val: 1}
	t.Root.Right.Left = &Node{Val: 0}
	t.Root.Right.Right = &Node{Val: 8}
	t.Root.Left.Right = &Node{Val: 2}
	t.Root.Left.Right.Left = &Node{Val: 7}
	t.Root.Left.Right.Right = &Node{Val: 4}
	t.Root.Left.Left = &Node{Val: 6}

	ans := findLCA(t.Root, 0, 3)
	if ans != nil {
		fmt.Println(ans.Val)
	} else {
		fmt.Println("Unable to compute")
	}

}

func findLCA(n *Node, v1 int, v2 int) *Node {
	if n == nil {
		return nil
	}

	if n.Val == v1 || n.Val == v2 {
		return n
	}

	leftCA := findLCA(n.Left, v1, v2)
	rightCA := findLCA(n.Right, v1, v2)

	// this is the case where LCA is found
	if leftCA != nil && rightCA != nil {
		return n
	}

	if leftCA != nil {
		// returning a non-nil value is the key here
		return leftCA
	}

	return rightCA
}
