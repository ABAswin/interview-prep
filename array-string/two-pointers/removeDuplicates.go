package main

func returnDistinct(arr []int) int {
	nextNonDuplicate := 1
	start := 0
	end := start + 1
	for end < len(arr)-1 {
		if arr[start] == arr[end] {
			end++
			continue
		}
		nextNonDuplicate++
		start = end
		end++
	}

	return nextNonDuplicate
}
