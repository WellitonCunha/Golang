package main

import "fmt"

func fibonacci(numero uint) uint {
	if numero <= 1 {
		return numero
	}

	return fibonacci(numero-2) + fibonacci(numero-1)
}

func main() {
	fmt.Println("Fibonacci")

	posicao := uint(12)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
