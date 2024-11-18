package main

// search - finds a pair in arr with a given targetSum.
func searchTargetSum(arr []int, targetSum int) []int {
	// TODO: Write your code here

	start := 0
	end := len(arr) - 1

	for start < end {
		if arr[start]+arr[end] == targetSum {
			return []int{start, end}
		}
		start++
		end--
	}

	return []int{-1, -1} // pair not found
}
