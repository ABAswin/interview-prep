/*
Input: [2, 5, 3, 10], target=30
Output: [2], [5], [2, 5], [3], [5, 3], [10]
Explanation: There are six contiguous subarrays whose product is less than the target.

algo:
contigous:
2,5,3,10

======================
works for non-contigous
===================
sort the array
2 , 3 , 5 , 10

for i

low,high:= 0, len(a)-1

	for low < high{
	 p = low*high

	if p < target{
		for i = low;i<=high;i++{
		s= append(s,[]int{a[low],a[high]})
		s=append(s,[]int{a[low]})
	}else{

high--}

}
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(productArrays([]int{2, 5, 3, 10}, 30))
	fmt.Println(productArrays([]int{8, 2, 6, 5}, 50))
}

// func productArraysNonContigous(a []int, target int) [][]int {
// 	res := make([][]int, 0)
// 	low, high := 0, len(a)-1
// 	for low < high {
// 		p := a[low] * a[high]
// 		if p < target {
// 			for i := high; i > low; i-- {
// 				res = append(res, []int{a[low], a[i]})
// 				res = append(res, []int{a[i]})
// 			}
// 		} else {
// 			high--
// 			continue
// 		}
// 		low++
// 	}

// 	return res
// }

func productArrays(a []int, target int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(a)-1; i++ {
		if a[i]*a[i+1] < target {
			res = append(res, []int{a[i], a[i+1]})
			res = append(res, []int{a[i]})
		} else if a[i] < target {
			res = append(res, []int{a[i]})
		}
	}
	if a[len(a)-1] < target {
		res = append(res, []int{a[len(a)-1]})
	}
	return res
}

// Post evaluation
// The above program does only for consecutive 2 elements , however there can be more than 2 contigous elements
// hence the more robust solution is to iterate accross all iterations and check the prodct O(n^2)
