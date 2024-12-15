/*
[7,5,8,9,3,1,4,6]
increasing order montonic stack
[1,3,4,5,6,7,8,9]
*/
package main

import "fmt"

type Stack struct {
	s   []int
	top int
}

func (st *Stack) push(val int) {
	if st.isEmpty() {
		st.s = append(st.s, val)
		st.top++
		return
	}
	temp := []int{}
	fmt.Println("the stack is", st.s, "and top is", st.top)
	for st.top != -1 && val > st.s[st.top] {
		temp = append(temp, st.pop())
	}
	st.s = append(st.s, val)
	st.top++
	for j := len(temp) - 1; j >= 0; j-- {
		st.s = append(st.s, temp[j])
		st.top++
	}

}

func (st *Stack) pop() int {
	if st.top == -1 {
		return -1
	}
	n := st.s[st.top]
	st.top--
	st.s = st.s[:len(st.s)-1]
	return n
}
func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func main() {
	var st Stack
	st.top = -1

	st.push(7)
	st.push(5)
	st.push(8)
	st.push(9)
	st.push(3)
	st.push(1)
	st.push(6)

	for st.top != -1 {
		fmt.Println(st.pop())
	}
}
