// Find the difference of scores when 2 players pick a number from start of array and if an even number is picked the array is reversed every time.
 
// [ 3,6,1,7,8,2,4,5]
package main

import "fmt"

func main() {
	fmt.Println(arrayGame([]int{3, 6, 1, 7, 8, 2, 4, 5}))
}
// TODO needs ro be improved
func arrayGame(a []int) int {

	p1Score := 0
	p2Score := 0
	left, right := 1, len(a)-1
	p1Score += a[0]
	front := true
	if a[0]%2 == 0 {
		front = !front
	}
	for i := 1; i < len(a); i++ {
		if front {
			if i%2 == 0 {
				p2Score += a[left]
			} else {
				p1Score += a[left]
			}
			left++
		} else {
			if i%2 == 0 {
				p2Score += a[right]
			} else {
				p1Score += a[right]
			}
			right--
		}
	}
	return p1Score - p2Score
}
