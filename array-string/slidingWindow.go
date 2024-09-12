//Given an array of positive integers nums and an integer k, find the length of the longest subarray whose sum is less than or equal to k.
package main

import "fmt"

func slidingWindow() int {
	// var a []int = []int{2, 3, 1, 4, 1, 2, 6}
	// k := 12
	a := []int{13, 14, 15} // All elements are greater than k
	k := 12
	left := 0
	right := 0
	max := 0
	sum := 0
	for right < len(a)-1 {
		sum += a[right]
		if sum <= k {
			right++
			max++
		} else {
			sum -= a[left]
			left++
			right++
		}
	}
	fmt.Println("The max is", max)
	return max
}
