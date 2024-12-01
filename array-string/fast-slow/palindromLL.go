/*
Input: 2 -> 4 -> 6 -> 4 -> 2 -> 2 -> null
Output: false

Input: 2 -> 4 -> 6 -> 4 -> 2 -> null
Output: true

find mid
reverse after mid
check palindrome
reverse to original
*/
package main

import (
	"fmt"
)

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
	fmt.Println(checkPalindrome(&ll))

	// ll := LL{}
	// ll.Head = &Node{Val: 1}
	// ll.Head.Next = &Node{Val: 2}
	// ll.Head.Next.Next = &Node{Val: 2}
	// ll.Head.Next.Next.Next = &Node{Val: 1}
	// fmt.Println(checkPalindrome(&ll))

	// ll := LL{}
	// ll.Head = &Node{Val: 1}
	// ll.Head.Next = &Node{Val: 2}
	// ll.Head.Next.Next = &Node{Val: 3}
	// ll.Head.Next.Next.Next = &Node{Val: 2}
	// ll.Head.Next.Next.Next.Next = &Node{Val: 1}
	// fmt.Println(checkPalindrome(&ll))

}

func checkPalindrome(ll *LL) bool {
	mid := findMid(ll)
	if mid != nil {
		fmt.Println("middle value", mid.Val)
	}
	firstHalf := ll.Head
	secondHalf := reverse(mid.Next)
	ispalindrome := true
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			ispalindrome = false
			break
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}
	mid.Next = reverse(mid.Next)
	return ispalindrome
}

func findMid(ll *LL) *Node {

	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		fmt.Println(slow.Val, fast.Val)
		slow = slow.Next
		fast = fast.Next.Next

	}
	return slow
}

func reverse(n *Node) *Node {
	var prev *Node
	temp := n

	for temp.Next != nil {
		curr := temp
		temp = temp.Next
		curr.Next = prev
		curr = prev
	}
	temp.Next = prev
	return temp
}
