package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Type string
}

func (c Circle) Area() float64 {
	fmt.Println("计算面积", c.Type)
	return 0
}

func (c Circle) Perimeter() float64 {
	fmt.Println("计算周长", c.Type)
	return 0
}

type Rectangle struct {
	Type string
}

func (r Rectangle) Area() float64 {
	fmt.Println("计算面积", r.Type)
	return 0
}

func (r Rectangle) Perimeter() float64 {
	fmt.Println("计算周长", r.Type)
	return 0
}

func main() {
	var c Circle
	c.Type = "Circle"
	
	var s1 Shape = c	
	s1.Area()
	s1.Perimeter()
	

	var r Rectangle
	r.Type = "Rectangle"
	var s2 Shape = r	
	s2.Area()
	s2.Perimeter()
}