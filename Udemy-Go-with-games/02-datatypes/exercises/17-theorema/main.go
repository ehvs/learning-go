package main

import (
	"fmt"
	"math"
)

// Start with:a2 + b2 = c2
// Put in what we know:5^2 + 12^2 = c2
// Calculate squares:25 + 144 = c2
// 25+144=169:169 = c2
// Swap sides:c2 = 169
// Square root of both sides:c = √169
// Calculate:c = 13

func main() {

	var a, b = 6, 7
	//var aFloat, aFloat = float64(a), float64(b)
	aFloat := float64(a)
	bFloat := float64(b)
	tmp := math.Sqrt(aFloat*aFloat + bFloat*bFloat)
	fmt.Printf("%.15f", tmp)
}
