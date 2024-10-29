/*
You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.
*/
package main

import "fmt"

func main() {

	cities := make([][]int, 5)
	fmt.Println(findCircleNum(cities))

}
func findCircleNum(isConnected [][]int) int {
	provinces := 0
	visited := make([]bool, len(isConnected))

	for cityNumber, _ := range isConnected {

		if !visited[cityNumber] {
			provinces++
			locateNeihbours(cityNumber, isConnected, visited)

		}
	}
	return provinces
}

func locateNeihbours(city int, isConnected [][]int, visited []bool) {
	visited[city] = true
	for i := 0; i < len(isConnected); i++ {
		if isConnected[city][i] == 1 && !visited[i] {
			locateNeihbours(i, isConnected, visited)
		}
	}
}
