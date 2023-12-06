package main

import "fmt"

func main() {
	// operadores

	soma := 1 + 1
	subtracao := 10 - 5
	divisao := 20 / 6
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// operadores lógicos
	fmt.Println(false && true)
	fmt.Println(true || false)
	fmt.Println(!false)
}
