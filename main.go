package main

import (
	"fmt"

	"github.com/jigth/go-algorithms/algorithms"
)

const (
	endFibo = 10
	endFact = 10
)

func main() {
	fmt.Println("Fibonacci:")
	for i := 0; i < endFibo; i++ {
		fibonacci := fmt.Sprintf("Fibo %d: %d", i, algorithms.Fibo(i))
		fmt.Println(fibonacci)
	}

	fmt.Println()
	fmt.Println("Factorial:")
	for i := 0; i < endFact; i++ {
		factorial := fmt.Sprintf("Fact %d: %d", i, algorithms.Fact(i))
		fmt.Println(factorial)
	}
}
