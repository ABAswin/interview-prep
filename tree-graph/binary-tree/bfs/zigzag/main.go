/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
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

type Queue []*Node

func main() {
	t := &BinaryTree{&Node{Val: 3}}
	t.Root.Left = &Node{Val: 9}
	t.Root.Right = &Node{Val: 20}
	t.Root.Right.Left = &Node{Val: 15}
	t.Root.Right.Right = &Node{Val: 7}

	var q Queue
	q = append(q, t.Root)
	fmt.Println(zigzag(q))

}

func zigzag(q Queue) [][]int {
	ret := make([][]int, 0)
	rev := false

	for len(q) > 0 {
		qSize := len(q)
		fmt.Println("levl sizes", qSize)
		t := make([]int, qSize)
		for i := 0; i < qSize; i++ {
			v := q[0]
			q = q[1:]
			fmt.Println("val", v.Val)

			if rev {
				t[len(t)-1-i] = v.Val
			} else {
				t[i] = v.Val
			}
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}

		ret = append(ret, t)
		rev = !rev
	}

	return ret

}
