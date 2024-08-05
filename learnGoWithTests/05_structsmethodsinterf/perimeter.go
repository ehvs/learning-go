package main

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

/*
	func Area(rectangle Rectangle) float64 {
		return rectangle.Width * rectangle.Height
	}
*/

/*
1. Construcao do struct
type <nome> struct

2. Construcao do metodo, similar a funcao a syntaxe.
func (<letraMin> <ReceiverType>) <nomedoMetodo>(<args>) <tipo>
*/
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

type Shape interface {
	Area() float64 // metodo
}

//Rectangle has a method called Area that
// returns a float64 so it satisfies the Shape interface
//Circle has a method called Area that returns
// a float64 so it satisfies the Shape interface
