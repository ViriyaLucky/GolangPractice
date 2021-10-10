// Exercise for https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
	"strconv"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Printf("iterasi ke "+strconv.Itoa(i)+" hasilnya %v\n", z)
	}
	return z
}

func main() {
	i := float64(169)
	fmt.Println(Sqrt(i), Sqrt(i) == math.Sqrt(i))
}
