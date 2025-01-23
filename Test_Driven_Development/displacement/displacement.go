/*
Let us assume the following formula for displacement s as a function of time t,acceleration a, initial velocity vo, and initial displacement so.
s = 1/2 a t^2 + v0 t + s0
This program first prompts the user to enter values for acceleration,
initial velocity, and initial displacement.
Then it prompts the user to enter a value for time and computes
the displacement after the entered time.
We define and use a function called GenDisplaceFn() which takes three float64
arguments:
  acceleration a,
  initial velocity v0,
  initial displacement s0.
GenDisplaceFn() returns a function which computes displacement as
a function of time, assuming the given values acceleration, initial velocity,
and initial displacement.
The function returned by GenDisplaceFn() takes one float64 argument t, r
epresenting time, and return one float64 argument which is the displacement
travelled after time t.
Example:
assume the following initial values:
  acceleration: a = 10,
  initial velocity: v0 = 2,
  initial displacement: s0 = 1.
The following statement can be used to call GenDisplaceFn()
to generate a function fn which will compute displacement as a function of time.
  fn := GenDisplaceFn(10, 2, 1)
Then to print the displacement after 3 seconds.
  fmt.Println(fn(3))
And hence the following statement to print the displacement after 5 seconds.
  fmt.Println(fn(5))
*/

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
