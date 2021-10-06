package main

import (
	"fmt"
	"unicode"
)

func main() {
	var X uint8 = 225
	fmt.Println(X, X-10)
	fmt.Println("hello world")
	var Z rune = 20

	fmt.Println(Z)
	fmt.Printf("\nThe type of Z is : %T\n", Z)
	a := 20.45
	b := 34.89

	// Subtraction of two
	// floating-point number
	c := b - a

	// Display the result
	fmt.Printf("Result is: %f", c)

	// Display the type of c variable
	fmt.Printf("\nThe type of c is : %T\n", c)

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	v1 := kvs["a"]

	fmt.Println(v1)
	v1 = "test"

	fmt.Println(v1)
	fmt.Println(kvs["a"])
	fmt.Println(SwapRune('a'))

}

func SwapRune(r rune) rune {
	switch {
	case 97 <= r && r <= 122:
		return r - 32
	case 65 <= r && r <= 90:
		return r + 32
	default:
		return r
	}
}
func SwapRune2(r rune) rune {
	if unicode.IsUpper(r) {
		r = unicode.ToLower(r)
	} else {
		r = unicode.ToUpper(r)
	}
	return r
}
