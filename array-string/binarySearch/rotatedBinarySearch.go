//[4,5,6,7,8,9,1,2] target = 9

/*
	if target > a[mid] &&
		if target >a[high]
		high = mid - 1
		else
		low = mid + 1
	if target < a[mid]
	 if target < a[low]
		low = mid + 1
		else
		high = mid - 1



*/

package main

import "fmt"

func main() {
	fmt.Println(rotatedBinarySearch([]int{4, 5, 6, 7, 8, 9, 1, 2}, 4))
}

func rotatedBinarySearch(a []int, target int) int {

	low, high := 0, len(a)-1
	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("low %d high %d mid %d\n", low, high, mid)
		if a[mid] == target {
			return mid
		}
		if a[low] <= a[mid] {
			if target < a[mid] && target >= a[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > a[mid] && target <= a[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		// if target > a[mid] {
		// 	if target > a[high] {
		// 		high = mid - 1
		// 	} else {
		// 		low = mid + 1
		// 	}
		// } else {
		// 	if target < a[low] {
		// 		low = mid + 1
		// 	} else {
		// 		high = mid - 1
		// 	}
		// }
	}

	return -1
}
