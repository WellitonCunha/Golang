package main

import "fmt"

func main() {
	numero := 10

	if numero >= 15 {
		fmt.Println("Maior ou igual a 15")
	} else {
		fmt.Println("Menor do que 15")
	}

	// variaveis criado dentro do scorpo do if não podem ser acessados fora
	if outroNumero := 1; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else if outroNumero > 5 {
		fmt.Println("Maior que cinco")
	} else {
		fmt.Println("Menor que zero")
	}
}
