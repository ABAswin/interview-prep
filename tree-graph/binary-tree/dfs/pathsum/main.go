//Given the root of a binary tree and an integer targetSum, return true if there exists a path from the root to a leaf such that the sum of the nodes on the path is equal to targetSum, and return false otherwise.

/*
				5
				/\
			4		8
			/		/\
		11		  13	4
		/\				  \
	   7  2					1
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

func main() {
	t := &BinaryTree{}
	t.Root = &Node{Val: 5}
	t.Root.Left = &Node{Val: 4}
	t.Root.Right = &Node{Val: 8}
	t.Root.Left.Left = &Node{Val: 11}
	t.Root.Left.Left.Left = &Node{Val: 7}
	t.Root.Left.Left.Right = &Node{Val: 2}
	t.Root.Right.Left = &Node{Val: 13}
	t.Root.Right.Right = &Node{Val: 4}
	t.Root.Right.Right.Right = &Node{Val: 1}

	fmt.Println(CheckTargetSum(t.Root, 0, 22))
}

func CheckTargetSum(n *Node, curr, target int) bool {
	if n == nil {
		return false
	}

	sum := curr + n.Val

	if n.Left == nil && n.Right == nil {
		return sum == target
	}

	left := CheckTargetSum(n.Left, sum, target)
	right := CheckTargetSum(n.Right, sum, target)

	return left || right
}
