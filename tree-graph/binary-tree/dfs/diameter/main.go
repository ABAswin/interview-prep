/*Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.*/

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
	t := &BinaryTree{Root: &Node{Val: 1}}
	// t.Root.Left = &Node{Val: 2}
	// t.Root.Left.Left = &Node{Val: 4}
	//t.Root.Left.Right = &Node{Val: 5}
	// t.Root.Left.Left.Left = &Node{Val: 5}
	t.Root.Right = &Node{Val: 2}
	t.Root.Right.Right = &Node{Val: 3}
	t.Root.Right.Right.Right = &Node{Val: 4}

	fmt.Println(calDiameter(t.Root))
}
func calDiameter(root *Node) int {
	var dfs func(root *Node, subMax int) int
	maxSum := 0
	dfs = func(root *Node, subMax int) int {

		if root == nil {
			return 0
		}

		left := dfs(root.Left, subMax)
		right := dfs(root.Right, subMax)

		subMax = max(left, right) + 1

		maxSum = max(maxSum, left+right)

		return subMax
	}
	dfs(root, 0)
	return maxSum
}

func max(a, b int) int {
	fmt.Println("Number compared", a, b)
	if a > b {
		return a
	}
	return b
}

//learning in this excercise experiment all the cases with both child ,one child , balanced tree, unbalanced tree.
