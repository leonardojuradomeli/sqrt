package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	y := float64(1)

	for i := 1; i < 10; i++ {
		fmt.Println("====== Iteracion ", i)
		fmt.Println("z = ", z)
		z *= z
		fmt.Println("z = ", z)
		z = z - x
		fmt.Println("z = ", z)
		z = z / (2 * y)
		fmt.Println("z = ", z)

		z = y - z
		y = z
		fmt.Println("====== Fin ", z, " \n ")
	}

	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
