/*
ip: [1,2,3,1] op: 4
ip: [2,7,9,3,1] op: 12
ip: [2,7,9,4,1,11,2,8,1] op: 15
*/
package main

import "fmt"

func main() {
	//a := []int{2, 7, 9, 4, 1, 11, 2, 8, 1}
	// a := []int{2, 7, 9, 3, 11,1,6}
	//a := []int{10, 2, 2, 12}
	a := []int{2, 7, 9, 4, 1, 11, 2, 8, 1}
	max := 0
	sum1 := 0
	sum2 := 0
	for i := 0; i <= len(a)-1; i += 2 {
		sum1 += a[i]
		if i != len(a)-1 {
			sum2 += a[i+1]
		}
	}
	if sum1 > sum2 {
		max = sum1
	} else {
		max = sum2
	}
	fmt.Println(max)
}
