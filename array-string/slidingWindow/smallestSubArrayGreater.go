/*
Input: arr = [2, 1, 5, 2, 8], S=7
Output: 1
Explanation: The smallest subarray with a sum greater than or equal to '7' is [8].
Example 3:

Input: arr = [3, 4, 1, 1, 6], S=8
Output: 3
Explanation: Smallest subarrays with a sum greater than or equal to '8' are [3, 4, 1] or [1, 1, 6].

Algo:

expand window as much to accomodate sum and calc min window size
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(smalledSubsetArray([]int{3, 4, 1, 1, 6}, 8))
	// fmt.Println(smalledSubsetArray([]int{2, 1, 5, 2, 8}, 7))
}

func smalledSubsetArray(a []int, target int) int {
	minWindow := len(a)
	sum := 0
	low := 0
	high := 0
	for high < len(a) {
		fmt.Println(sum)
		sum += a[high]
		for sum >= target {
			minWindow = min(minWindow, high-low+1)
			sum -= a[low]
			if low < len(a) {
				low++
			}
		}
		high++
	}

	return minWindow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
