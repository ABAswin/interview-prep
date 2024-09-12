// package main

// import "fmt"

// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// }

// type BinaryTree struct {
// 	Root *Node
// 	Len  int
// }

// func (t *BinaryTree) Insert(val int) {
// 	n := Node{
// 		Val: val,
// 	}
// 	if t.Root == nil {
// 		t.Root = &n
// 		return
// 	}

// 	insertNode(t.Root, &n)

// }
// func insertNode(root *Node, node *Node) {

// 	if node.Val < root.Val {
// 		if root.Left != nil {
// 			insertNode(root.Left, node)
// 		} else {
// 			root.Left = node
// 		}
// 	} else if node.Val > root.Val {
// 		if root.Right != nil {
// 			insertNode(root.Right, node)
// 		} else {
// 			root.Right = node
// 		}
// 	}
// }

// func (t *BinaryTree) inorderTraversal() {
// 	inorderNodeTraversal(t.Root)
// }

// func inorderNodeTraversal(node *Node) {
// 	if node != nil {
// 		inorderNodeTraversal(node.Left)
// 		fmt.Println(node.Val)
// 		inorderNodeTraversal(node.Right)
// 	}
// }

// func (t *BinaryTree) Delete(val int) int {
// 	// deleteNode(t.Root, val)
// }

// func deleteNode(val int)

// func main() {
// 	tree := BinaryTree{}
// 	tree.Insert(5)
// 	tree.inorderTraversal()
// 	tree.Insert(3)
// 	tree.inorderTraversal()

// }
