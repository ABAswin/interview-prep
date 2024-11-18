// [1,2,4,5,6,7,8,9,11] target = 9
/*
Algo:
	mid = (low+high)/2
	if a[mid] == target {
	return index }
if target > a[mid]
	low = mid +1
else high = mid -1

*/
package main

// func main() {
// 	fmt.Println(binarySearch([]int{1, 2, 4, 5, 6, 7, 8, 9}, 4))

// }

func binarySearch(a []int, target int) int {
	low, high := 0, len(a)-1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == target {
			return mid
		}
		if target > a[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
