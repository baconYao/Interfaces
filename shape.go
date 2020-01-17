package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base float64
	height float64
}

type square struct {
	sideLength float64
}

func printArean(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}