package main

import (
	"fmt"
	"math"
)

const (
	Precision          float64 = 1.0e-14
	NumberOfIterations int     = 100
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 1; math.Abs((z*z-x)/(2*z)) > Precision && i <= NumberOfIterations; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration: %d, Square root aprox: %f\n", i, z)
	}
	return z
}

func main() {
	var number float64 = 63
	fmt.Printf("Square root aprox of %f: %f\n", number, Sqrt(number))
}
