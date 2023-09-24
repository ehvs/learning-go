package main

import (
	"fmt"
)

// In general, x and y must satisfy (x)² + (y)² < radius².
// center of circle C (0,0)

func main() {
	radius := 4
	x1, y1 := -2, 1
	result := ((x1 * x1) + (y1 * y1)) <= (radius * radius)
	fmt.Println(x1, ",", y1, "->", result)

	x2, y2 := 1, 4
	result = (x2*x2)+(y2*y2) <= (radius * radius)
	fmt.Println(x2, ",", y2, "->", result)

	x3, y3 := 2, 2
	result = (x3*x3)+(y3*y3) <= (radius * radius)
	fmt.Println(x3, ",", y3, "->", result)

}
