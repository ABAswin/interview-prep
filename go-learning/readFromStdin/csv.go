/*
"The ABC of Go";25.5;1500
"Functional Programming with Go";56;280
"Go for It";45.9;356
"The Go Way";55;500
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float32
	quantity int
}

var BookStore []Book

// func main() {
// 	f, err := os.Open("data.csv")
// 	if err != nil {
// 		fmt.Println("error opening file", err)
// 	}
// 	rd := csv.NewReader(f)

// 	records, err := rd.ReadAll()
// 	if err != nil {
// 		fmt.Println("error reading CSV file ", err)
// 		return
// 	}
// 	for _, r := range records {
// 		pr, _ := strconv.ParseFloat(r[1], 32)
// 		qa, _ := strconv.Atoi(r[2])

// 		BookStore = append(BookStore,
// 			Book{title: r[0],
// 				price:    float32(pr),
// 				quantity: qa,
// 			})
// 	}
// 	fmt.Println("records", BookStore)
// }

func main() {

	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := bufio.NewReader(f)

	for {
		data, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		fields := strings.Split(data, ";")
		pr, _ := strconv.ParseFloat(fields[1], 32)
		qa, _ := strconv.Atoi(fields[2])
		BookStore = append(BookStore, Book{
			title:    fields[0],
			price:    float32(pr),
			quantity: qa,
		})
	}
	fmt.Println(BookStore)
}
