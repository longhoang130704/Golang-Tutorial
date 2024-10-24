package main

import (
	"fmt"
	"math"

	"example.com/greetings"
	"rsc.io/quote"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v, "co the su dung o tat ca cac else cua if")
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}


func swap(a, b string) (string, string) {
	return b, a;
}

func main() {
    fmt.Println(quote.Hello());
	message := greetings.Hello("Long");
	fmt.Println(message);
	fmt.Println(math.Pi);
	a, b := swap("Hello", "World");
	fmt.Println(a, b);
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
