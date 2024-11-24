package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// // approach during test
// func main() {
// 	var str1Size, str2Size int
// 	var str1, str2 string

// 	_, err := fmt.Scanf(`%d
// 	%s
// 	%d
// 	%s`, &str1Size, &str1, &str2Size, &str2)
// 	if err != nil {
// 		fmt.Printf("error in scanf %v\n", err)
// 	}
// 	a := strings.Split(str1, ",")
// 	b := strings.Split(str2, ",")
// 	fmt.Println(str1Size, str2Size, a, b)
// 	if str1Size != str2Size {
// 		fmt.Println("NO")
// 		return
// 	}
// 	fmt.Println(checkAlmostEqual(str1Size, a, b))

// }

func main() {

	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	str1Size, _ := strconv.Atoi(input.Text())
	input.Scan()
	str1 := strings.Split(input.Text(), ",")
	input.Scan()
	str2Size, _ := strconv.Atoi(input.Text())
	input.Scan()
	str2 := strings.Split(input.Text(), ",")
	if str1Size != str2Size {
		fmt.Println("NO")
		return
	}
	fmt.Println(checkAlmostEqual(str1Size, str1, str2))

}

func checkAlmostEqual(numberOfStrings int, a []string, b []string) string {
	for k := 0; k < numberOfStrings; k++ {
		if len(a[k]) != len(b[k]) {
			return "NO"
		}
		ma := make(map[string]int, len(a[k]))
		mb := make(map[string]int, len(b[k]))
		for i := 0; i < len(a[k]); i++ {
			ma[string(a[k][i])]++
			mb[string(b[k][i])]++
		}
		fmt.Println("values of ma,mb", ma, mb)
		for k, v := range ma {
			if abs(mb[k]-v) > 3 {
				return "NO"
			} else {
				delete(mb, k)
			}
		}
		for _, n := range mb {
			if n > 3 {
				return "NO"
			}
		}
		fmt.Println("value of k", k)
	}

	return "YES"
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
