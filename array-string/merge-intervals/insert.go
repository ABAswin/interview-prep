/*
Given a list of non-overlapping intervals sorted by their start time, insert a given interval at the correct position and merge all necessary intervals to produce a list that has only mutually exclusive intervals.

Example 1:

Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]
Output: [[1,3], [4,7], [8,12]]
Explanation: After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].

Algo:
insert the interval and merge intervals
*/
package main

import "fmt"

func main() {
	fmt.Println(insertInterval([][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 6}))

}

func insertInterval(a [][]int, b []int) [][]int {
	var final [][]int
	var i int
	for i < len(a) && b[0] > a[i][0] {
		final = append(final, a[i])
		i++
	}
	fmt.Println(final)

	for i < len(a) && a[i][0] <= b[1] {

		b[0] = min(a[i][0], b[0])
		b[1] = max(a[i][1], b[1])
		final = append(final, []int{b[0], b[1]})
		i++

	}
	fmt.Println(final)
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	fmt.Println(final)
	return final
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
