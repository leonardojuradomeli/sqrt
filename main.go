package main

import (
	"fmt"
)

func sqrtLogic() func(float64) float64 {
	z := float64(1)

	return func(x float64) float64 {
		// Resolve next equation
		// z - (z^2 - x)/2.z
		y := z
		z *= z
		z = z - x
		z = z / (2 * y)
		z = y - z
		y = z
		return z
	}
}

func sqrt(x float64) float64 {
	sqrt := sqrtLogic()

	var result float64
	i := 0

	next := sqrt(x)
	for next != result && i < 20 {
		//fmt.Println("Next :", next, " result :", result)
		result = next
		next = sqrt(x)
		i++
	}

	fmt.Println("This result was generate in", i, "attempts")
	return result
}

func main() {
	in := float64(100)
	fmt.Println("Square root of ", in, " is ", sqrt(in))
}
