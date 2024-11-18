//Given an integer array which is rotated at some index and a target, find the index of the target in the array or return -1 if target is not found.

// a := [4,5,1,2,3]

// low = 0, high = 4 - a[low] > a[high] -> low = low+high/2
// low = 2, high = 4 - a[low] < a[high] , a[low] < target , a[high] > target

//[4,5,6,7,0,1,2] target = 0

// brute force approach - iterating over the array - O(n)
// optimised apprach  - Binary search - O(log n)

package main

// func main() {

// 	fmt.Println(findElementRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0))

// }

func findElementRotated(a []int, target int) int {
	low := 0
	high := len(a) - 1

	for {
		if a[low] == target {
			return low
		} else if a[high] == target {
			return high
		}

		if a[low] > a[high] {
			low = (low + high) / 2
		} else {
			if target < a[low] {
				low--
			} else {
				low++
			}

		}

	}
}
