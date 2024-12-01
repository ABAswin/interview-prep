package main

import(
  "fmt"
  )

type Square struct {
	side float32
}

type Triangle struct {	// implement this struct
	base float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface { // implement this interface 
	Perimeter() float32
}
func main() {
var pi PeriInterface
sq := &Square{side: 5}
pi = sq
fmt.Println("the perimeter ",pi.Perimeter())	
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 { // implement method called on square to calculate its perimeter
	return 4* sq.side
}

func (tr *Triangle) Area() float32 { // implement method called on triangle to calculate its area
	return 0.5* tr.base * tr.height
}
