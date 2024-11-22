/*
Input: str1="xp#", str2="xyz##"
Output: true
Explanation: After applying backspaces the strings become "x" and "x" respectively.
In "xyz##", the first '#' removes the character 'z' and the second '#' removes the character 'y'.

Input: str1="xywrrmp", str2="xywrrmu#p"
Output: true
Explanation: After applying backspaces the strings become "xywrrmp" and "xywrrmp" respectively.

Algo:

high = len(a)-1

	for high >0 {
		if a[high] != b[]
	}
*/
package main

import "fmt"

func main() {
	fmt.Println(backspace("xp#", "xyz##"))
}

func backspace(a, b string) bool {

	lenb := len(b) - 1
	lena := len(a) - 1
	diffa, diffb := 0, 0
	for lenb > 0 && lena > 0 {
		if a[lena] != b[lenb] {
			return false
		} else if a[lena] == b[lenb] {
			lena -= diffa
			lenb -= diffb
			diffa, diffb = 0, 0
		}
		if a[lena] == '#' {
			diffa++
		}
		if b[lenb] == '#' {
			diffb++
		}
		lena--
		lenb--
	}
	return true
}
