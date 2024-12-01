package main

// type Node struct {
// 	Val  int
// 	Next *Node
// }

// type LL struct {
// 	Head *Node
// }

/*
1 2 3 4
2

11
23
32
44

1 2 3 4 5
3

11
23
35
42
54
32
44

1 2 3 4 5
2

11
23
35
42
54
22

14
25
32
43
52
23
34
45
52
23
32

1 2 3 4 5 6 7 8
3
guess - 3
11
23
35
47
53
65
77
guess not right

slow and fast meet somewhere not any constant point
*/
// func main() {
// 	ll := LL{&Node{Val: 1}}
// 	ll.Head.Next = &Node{Val: 2}
// 	ll.Head.Next.Next = &Node{Val: 3}
// 	ll.Head.Next.Next.Next = &Node{Val: 4}
// 	// 4 is connected to 2
// 	ll.Head.Next.Next.Next.Next = ll.Head.Next
// 	fmt.Println(findCycleStart(ll).Val)
// }

func findCycleStart(ll LL) *Node {
	slow := ll.Head
	fast := ll.Head
	var count int
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			count = findCycleCount(slow)
			break
		}
	}
	return findStart(ll.Head, count)

}

func findCycleCount(n *Node) int {
	temp := n
	count := 0
	for temp.Next != n {
		temp = temp.Next
		count++
	}
	return count
}

func findStart(head *Node, count int) *Node {
	ptr1 := head
	ptr2 := head

	for count >= 0 {
		ptr2 = ptr2.Next
		count--
	}
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}
