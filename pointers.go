package main

import (
	"fmt"

	"math"
)

var P float64 = 35
var r float64
var R *float64

func main() {
	pi := math.Pi

	r = P / (2.0 * pi)
	R = &r
	circleArea := pi * (*R * (*R))
	fmt.Println(circleArea)

}
