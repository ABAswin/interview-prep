/*
Input: 23
Output: true (23 is a happy number)
Explanations: Here are the steps to find out that 23 is a happy number:

 = 4 + 9 = 13
 = 1 + 9 = 10
 = 1 + 0 = 1

 split
 2,3
 map


 10201923

*/

package main

import "fmt"

// func main() {
// 	//fmt.Println(isMagic(23))
// 	fmt.Println(isMagic(12))
// }

func isMagic(n int) bool {
	ma := make(map[int]int, 0)
	for {
		num := checkMagic(n)
		fmt.Println(num)
		if num == 1 {
			return true
		}
		if _, ok := ma[num]; ok {
			return false
		}
		ma[num]++
		n = num
	}
}

func checkMagic(n int) int {
	var a []int
	for {
		if n == 0 {
			break
		}
		num := n % 10
		a = append(a, num)
		n = n / 10
	}
	fmt.Println(a)
	sum := 0
	// can be further optimised by not creating array and directly adding to sum
	for _, v := range a {
		sum += v * v
	}
	return sum
}
