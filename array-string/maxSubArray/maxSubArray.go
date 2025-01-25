/*
[-1,4,-9,10]

[-1,4,32,1,-20,5,7,1,4,8,0]
[2,1,-1,-3,4]
[-2,-1,-4,1,-5]
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-1, 4, -9, 10}))
	fmt.Println(maxSubArray([]int{2, 1, -1, -3, 4}))
	fmt.Println(maxSubArray([]int{-2, -1, -4, 1, -5}))
}

func maxSubArray(a []int) int {
	max := -math.MaxInt16
	indMax := -math.MaxInt16
	currSum := 0

	for _, v := range a {
		currSum += v

		if currSum > max {
			max = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
		if v > indMax {
			indMax = v
		}
	}
	if max <= 0 && indMax < 0 {
		return indMax
	}
	return max
}
