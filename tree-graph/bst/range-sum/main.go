/*Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high]. */

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
	t := &BinaryTree{Root: &Node{Val: 20}}
	t.Root.Left = &Node{Val: 15}
	t.Root.Left.Left = &Node{Val: 10}
	t.Root.Left.Right = &Node{Val: 18}
	t.Root.Right = &Node{Val: 25}
	t.Root.Right.Right = &Node{Val: 30}
	fmt.Println(rangeSum(t.Root, 16, 26))
}
func rangeSum(n *Node, low, high int) int {
	if n == nil {
		return 0
	}
	sum := 0

	fmt.Println("Node value ", n.Val)
	if n.Val > low {
		sum = rangeSum(n.Left, low, high)
	}
	if n.Val < high {
		sum = rangeSum(n.Right, low, high)
	}
	if n.Val >= low && n.Val <= high {
		fmt.Println("Add Node value ", n.Val)
		sum += n.Val
	}
	fmt.Println("returning value", sum)
	return sum
}
