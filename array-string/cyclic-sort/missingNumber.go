/*
Example 1:

Input: [4, 3, 1, 0]
Output: 2
Example 2:

Input: [8, 3, 5, 2, 4, 6, 0, 1]
Output: 7
Constraints:

n == nums.length
1 <= n <=
0 <= nums[i] <= n
All the numbers of nums are unique.
*/
package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{4, 3, 1, 0}))
	fmt.Println(missingNumber([]int{8, 3, 5, 2, 4, 6, 0, 1}))
}

func missingNumber(a []int) int {
	i := 0
	for i < len(a) {
		if a[i] < len(a) && a[i] != a[a[i]] {
			a[i], a[a[i]] = a[a[i]], a[i]
		} else {
			i++
		}
	}
	for i, v := range a {
		if i != v {
			return i
		}
	}

	return i
}
