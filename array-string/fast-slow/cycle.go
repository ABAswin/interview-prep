// find out if a linked list is cycle or not
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LL struct {
	Head *Node
}

func main() {

	ll := &LL{Head: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: nil}}}}
	fmt.Println(isCycle(ll.Head))

}

func isCycle(node *Node) bool {

	slow := node
	fast := node
	for fast != nil && fast.Next != nil {
		slow = node.Next
		fast = node.Next.Next
	}
	if slow == fast {
		return true
	}
	return false
}
