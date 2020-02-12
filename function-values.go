/*
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.

*/
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func process(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	mul := func(x, y float64) float64 {
		return x * y
	}

	fmt.Println(mul(2, 3))
	fmt.Println(process(mul))
}
