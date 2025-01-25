package main

import (
	"sort"
)

func main() {

	// fmt.Println(findDuplicate([]int{1, 2, 3, 4}))
	// fmt.Println(findDuplicate([]int{1, 2, 3, 2}))

	// fmt.Println(findAnagrams("jar", "jam"))
	// fmt.Println(findAnagrams("racecar", "carrace"))

	// fmt.Println(twoSum([]int{3, 4, 5, 6}, 7))
	// fmt.Println(twoSum([]int{2, 3, 4, 5}, 7))
	// fmt.Println(twoSum([]int{5, 5}, 10))
	// fmt.Println(twoSum([]int{5, 5}, 7))
	// fmt.Println(twoSum([]int{4, 5, 6}, 10))

	// fmt.Println(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))

	// fmt.Println(topKFrequentElements([]int{1, 2, 2, 3, 3, 3}, 2))
	// fmt.Println(topKFrequentElements([]int{7, 7}, 1))

}

func findDuplicate(a []int) bool {
	m := make(map[int]int)
	for _, v := range a {
		if _, ok := m[v]; ok {
			return false
		}
		m[v]++
	}
	return true
}

func findAnagrams(a, b string) bool {

	m := make(map[rune]int)
	for i := 0; i < len(a); i++ {
		m[rune(a[i])]++
		m[rune(b[i])]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

/*
Input:
nums = [3,4,5,6], target = 7

Output: [0,1]

*/

func twoSum(a []int, target int) []int {
	left, right := 0, len(a)-1
	for left < right {
		sum := a[left] + a[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}

/*
	Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
*/
func groupAnagrams(a []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for i := 0; i < len(a); i++ {
		temp := a[i]
		runes := []rune(temp)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		if _, ok := m[string(runes)]; ok {
			m[string(runes)] = append(m[string(runes)], temp)
		} else {
			m[string(runes)] = []string{temp}
		}

	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

/*
	Input: nums = [1,2,2,3,3,3], k = 2

Output: [2,3]
*/
func topKFrequentElements(a []int, k int) []int {
	var res []int
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
		// if val, _ := m[v]; val >= k {
		// 	res = append(res, v)
		// }
	}
	for key, v := range m {
		if v >= k {
			res = append(res, key)
		}
	}
	return res
}
