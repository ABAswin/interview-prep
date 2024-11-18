package main

import (
	"sort"
)

// func main() {
// 	//searchTargetSum([]int{1, 2, 4, 6, 9}, 5)
// 	// fmt.Println(returnDistinct([]int{2, 2, 2, 11}))
// 	fmt.Println(smallestTriplets([]int{-2, 0, 1, 3}, 2))
// }

func smallestTriplets(arr []int, target int) int {
	count := 0
	sort.Ints(arr)

	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum < target {
				count = count + right - left
				left++
			} else {
				right--
			}
		}
	}
	return count
}
