package main

import "fmt"

func main() {

	c := make(chan int)
	go receive(c)
	for i := 0; i < 5; i++ {
		c <- i
	}

}

func receive(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
