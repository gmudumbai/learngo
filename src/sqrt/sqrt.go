package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	const DELTA = 1e-10
	const INITIAL_Z = 100.0

	z = INITIAL_Z

	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}

	for zz := step(); math.Abs(zz-z) > DELTA; {
		z = zz
		zz = step()
	}
	
	return
}

func main() {
	fmt.Println(Sqrt(500))
	fmt.Println(math.Sqrt(500))
}