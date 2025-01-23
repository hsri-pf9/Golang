package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64
	fmt.Println("Enter the value of acceleration, initial velocity, initial displacement")
	fmt.Scanln(&a, &v0, &s0)

	dispFn := GenDisplaceFn(a, v0, s0)

	fmt.Println("Enter time")
	fmt.Scan(&t)

	displacement := dispFn(t)
	fmt.Printf("Displacement after %0.2f seconds: %0.2f m\n", t, displacement)
}
