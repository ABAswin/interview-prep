package main

import (
	"strings"
)

var nrchars, nrwords, nrlines int

// func main() {
// 	read := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the input, Press S and enter to end the input and obtain stats")
// 	for {
// 		input, err := read.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("Error reading input", err)
// 			os.Exit(1)
// 		}
// 		if input == "S\n" {
// 			fmt.Println("Stats:")
// 			fmt.Println("Chars:", nrchars)
// 			fmt.Println("words:", nrwords)
// 			fmt.Println("lines:", nrlines)
// 			os.Exit(0)
// 		}
// 		counters(input)
// 	}

// }

func counters(input string) {
	nrchars += len(input) - 2
	nrwords += strings.Count(input, " ") + 1
	nrlines++
}
