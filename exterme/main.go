/*
https://go.dev/play/p/hs4rYmOTAU8
linked list problem
explain interfaces with a employee example
*/
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func main() {
	ll := LinkedList{Head: &Node{Val: 1}}
	ll.Head.Next = &Node{Val: 2}
	ll.Head.Next.Next = &Node{Val: 3}
	ll.Head.Next.Next.Next = &Node{Val: 4}
	print(ll.Head)
	reverse(ll.Head)

}

func reverse(head *Node) {

	var prev *Node
	curr, temp := head, head

	for temp.Next != nil {
		temp = temp.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	temp.Next = prev
	head = temp

	print(head)
}

func print(n *Node) {
	for temp := n; temp != nil; {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
