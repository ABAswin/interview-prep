/*
Input: [4, 1, 2, -1, 1, -3], target=1
Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
Explanation: Both the quadruplets add up to the target.

sort
[-3,-1,1,1,2,4]
*/
package main

import (
	"sort"
)

// func main() {

// 	// fmt.Println(quadruplets([]int{4, 1, 2, -1, 1, -3}, 1))
// 	fmt.Println(quadruplets([]int{1, 1, 1, 1, 1, 1}, 4))
// }

func quadruplets(arr []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(arr)
	for i := 0; i < len(arr)-3; i++ {

		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr)-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			low, high := j+1, len(arr)-1
			for low < high {
				sum := arr[i] + arr[j] + arr[low] + arr[high]
				if sum == target {
					res = append(res, []int{arr[i], arr[j], arr[low], arr[high]})
					low++
					high--
					for low < high && arr[low] == arr[low+1] {
						low++
					}
					for low < high && arr[high] == arr[high-1] {
						high--
					}

				} else if sum < target {
					low++
				} else {
					high--
				}

			}
		}
	}

	return res
}
