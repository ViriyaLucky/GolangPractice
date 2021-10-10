package main

import "fmt"

func adder() func(string, ...int) int {
	sum := 0
	return func(x string, y ...int) int {
		for _, value := range y {
			sum += value
		}
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos("test", i),
			neg("negatif", -2*i),
		)
	}
}
