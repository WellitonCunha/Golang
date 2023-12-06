package main

import "fmt"

func main() {

	var valor int
	var ponteiro *int

	valor = 10
	ponteiro = &valor
	fmt.Println("valor: ", valor)
	fmt.Println("endereço: ", *ponteiro)
}
