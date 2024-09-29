/*Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Input: root = [3,9,20,null,null,15,7]
Output: 2
Example 2:

Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5
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

/*
		3
		/\
	   5   	1
	  / \  	/\
	6	 2  0  8
		/\
		7 4





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
	t.Root.Left.Left.Left = &Node{Val: 10}
	// t.Root = &Node{Val: 2}
	// t.Root.Right = &Node{Val: 3}
	// t.Root.Right.Right = &Node{Val: 4}
	// t.Root.Right.Right.Right = &Node{Val: 5}

	fmt.Println(minDepth(t.Root))
}

func minDepth(n *Node) int {
	if n == nil {
		return 0
	}

	if n.Left == nil && n.Right == nil {
		return 1
	}
	left, right := 0, 0
	left += minDepth(n.Left)
	right += minDepth(n.Right)

	if left < right && left != 0 {
		return left + 1
	} else if right != 0 {
		return right + 1
	}
	return 1
}
