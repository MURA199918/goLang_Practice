package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Displacement calculation program.")
	a := 0.0
	v := 0.0
	s := 0.0
	t := 0.0
	fmt.Println("Enter acceleration value: ")
	fmt.Scanln(&a)
	fmt.Println("Enter velocity: ")
	fmt.Scanln(&v)
	fmt.Println("Enter initial distance: ")
	fmt.Scanln(&s)
	function1 := GenDisplaceFn(a, v, s)
	fmt.Println("Enter time: ")
	fmt.Scanln(&t)
	fmt.Println("Displacement for given time: ", function1(t))
}

//GenDisplaceFn function to find displacement
func GenDisplaceFn(acc, velocity, initialDistance float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return ((initialDistance) + (velocity * time) + (0.5 * acc * (math.Pow(time, 2))))
	}
	return fn
}
