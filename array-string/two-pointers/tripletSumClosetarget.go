package main

import (
	"math"
	"sort"
)

//-1, 2, 1, -4
// -4,-1,1,2

func tripletSumCloseTarget(arr []int, targetSum int) int {
	closestSum := math.MaxInt
	sort.Ints(arr)

	for idle := 0; idle < len(arr)-2; idle++ {
		left := idle + 1
		right := len(arr) - 1
		for left < right {
			sum := arr[idle] + arr[left] + arr[right]
			if abs(targetSum-sum) < abs(targetSum-closestSum) {
				closestSum = sum
			}
			if sum < targetSum {
				left++
			} else {
				right--
			}
		}

	}

	return closestSum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
