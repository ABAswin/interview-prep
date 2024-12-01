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
	ll := LL{}
	ll.Head = &Node{Val: 1}
	ll.Head.Next = &Node{Val: 2}
	ll.Head.Next.Next = &Node{Val: 3}
	ll.Head.Next.Next.Next = &Node{Val: 4}
	_ = reverse(&ll)
	//print(list)
}

func reverse(ll *LL) *LL {

	var prev *Node
	temp := ll.Head
	for temp.Next != nil {
		curr := temp
		//fmt.Println(temp.Val)
		temp = temp.Next
		curr.Next = prev
		prev = curr
	}
	temp.Next = prev
	ll.Head = temp
	print(ll.Head)
	return ll
}

func print(ll *Node) {
	temp := ll
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}

}
