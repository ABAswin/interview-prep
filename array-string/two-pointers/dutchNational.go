/*
Input: [2, 2, 0, 1, 2, 0]
Output: [0 0 1 2 2 2 ]

0,1,2

each of 0,1,2 have their own pointers



so if 0 is found increment i and low is incremented as 0 is found
if 0:
	swap(a[i],a[low])
	low++
	i++
if 1:
	i++
if 2:
	swap(a[i],a[high])
 	high--
	i++

	2, 2, 0, 1, 2, 0
	0,2,0,1,2,2

	1,2,1,0,1,0,2 - l,h,i = 0,6,0
	1,2,1,0,1,0,2 - l,h,i = 0,6,1
	1,2,1,0,1,0,2 - l,h,i = 0,5,2
	1,2,1,0,1,0,2 - l,h,i = 0,5,3
	0,2,1,1,1,0,2 - l,h,i = 1,5,4
	0,2,1,1,1,0,2 - l,h,i = 1,5,5
	0,0,1,1,1,2,2 - l,h,i = 2,5,6


*/

package main

// func main() {

// 	fmt.Println(dutchNationalFlag([]int{2, 2, 0, 1, 2, 0}))
// }

func dutchNationalFlag(a []int) []int {

	low, high, i := 0, len(a)-1, 0
	for i <= high {
		if a[i] == 0 {
			a[i], a[low] = a[low], a[i]
			low++
			i++
		} else if a[i] == 2 {
			a[i], a[high] = a[high], a[i]
			high--

		} else {
			i++
		}

	}
	return a
}
