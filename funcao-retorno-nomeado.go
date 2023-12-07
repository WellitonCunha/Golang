package main

import "fmt"

func calcularDoisValores(valor1, valor2 int) (somar int, subtrair int) {
	somar = valor1 + valor2
	subtrair = valor1 - valor2
	return
}

func main() {
	fmt.Println("Funções nomeadas")

	fmt.Println(calcularDoisValores(10, 7))
}
