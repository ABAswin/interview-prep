// The below solution is correct however is O(n*m) hence consider every time occurance as events and the time max cumulative value would be the ans
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	start := []int{1, 2, 1, 3}
	end := []int{5, 3, 3, 4}
	// start := []int{1, 1, 1, 1}
	// end := []int{5, 5, 5, 5}
	// start := []int{1, 2, 5}
	// end := []int{4, 5, 7}

	fmt.Println(maxOccuranceOptimised(start, end))
}

func maxOccurance(start, end []int) int {
	t := make(map[int]int, len(start))

	for i := 0; i < len(start); i++ {

		s := start[i]
		e := end[i]
		for j := s; j <= e; j++ {
			t[j]++
		}
	}
	maxValue, minVal := 0, math.MaxInt
	for i, v := range t {

		if v > maxValue {
			maxValue = v
			minVal = i

		} else if v == maxValue && i < minVal {
			minVal = i
		}
	}
	return minVal
}

func maxOccuranceOptimised(start, end []int) int {
	t := make(map[int]int, 0)

	for i := 0; i < len(end); i++ {
		t[start[i]]++
		t[end[i]+1]--

	}
	fmt.Println("The formed map", t)
	var times []int
	for k := range t {
		times = append(times, k)
	}
	sort.Ints(times)
	fmt.Println("Sorted keys", times)

	currConn, maxConn, minTime := 0, 0, 0

	for i := range times {
		if currConn == t[i] {
			currConn += t[i]
			continue
		}
		currConn += t[i]
		if currConn > maxConn {
			maxConn = currConn
			minTime = i
		}
	}
	return minTime
}
