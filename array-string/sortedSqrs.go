package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{-4, -1, 0, 3, 10}
	// a := []int{-10, 5, 1, 2, 4, 7}
	fmt.Println(a)
	fmt.Println(sortedSquares(a))
}

func sortedSquares(arr []int) []int {
	temp := make([]int, len(arr))
	i := 0
	j := len(arr) - 1
	for itr := len(arr) - 1; itr >= 0; itr-- {
		if math.Abs(float64(arr[i])) < math.Abs(float64(arr[j])) {
			temp[itr] = arr[j] * arr[j]
			j--
		} else {
			temp[itr] = arr[i] * arr[i]
			i++
		}
	}
	return temp
}
