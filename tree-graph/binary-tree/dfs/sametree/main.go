//Given the roots of two binary trees p and q, check if they are the same tree. Two binary trees are the same tree if they are structurally identical and the nodes have the same values.

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
	t1 := &BinaryTree{}
	t1.Root = &Node{Val: 2}
	t1.Root.Left = &Node{Val: 5}
	t1.Root.Right = &Node{Val: 3}
	t1.Root.Left.Left = &Node{Val: 7}
	t1.Root.Left.Right = &Node{Val: 8}
	t1.Root.Left.Right.Right = &Node{Val: 9}

	t2 := &BinaryTree{}
	t2.Root = &Node{Val: 2}
	t2.Root.Left = &Node{Val: 5}
	t2.Root.Right = &Node{Val: 3}
	t2.Root.Left.Left = &Node{Val: 7}
	t2.Root.Left.Right = &Node{Val: 8}
	//t2.Root.Left.Right.Right = &Node{Val: 9}

	fmt.Println(checkSame(t1.Root, t2.Root))
}
func checkSame(n1 *Node, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.Val != n2.Val {
		return false
	}

	left := checkSame(n1.Left, n2.Left)
	right := checkSame(n1.Right, n2.Right)

	return left && right

}
