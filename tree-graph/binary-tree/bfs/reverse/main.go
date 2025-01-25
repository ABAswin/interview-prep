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
	var bt BinaryTree
	bt.Root = &Node{Val: 1}
	bt.Root.Left = &Node{Val: 2}
	bt.Root.Left.Left = &Node{Val: 4}
	bt.Root.Left.Right = &Node{Val: 5}
	bt.Root.Right = &Node{Val: 3}
	bt.Root.Right.Left = &Node{Val: 6}
	bt.Root.Right.Right = &Node{Val: 7}

	a := reverse(bt.Root)

	for _, v := range a {
		fmt.Println(v)
	}
}

func reverse(root *Node) [][]int {
	q := []*Node{root}
	final := [][]int{}
	for len(q) != 0 {
		n := len(q)
		result := []int{}
		for i := 0; i < n; i++ {
			v := q[0]
			q = q[1:]
			result = append(result, v.Val)
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		final = append([][]int{result}, final...)
	}
	return final
}
