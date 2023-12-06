package main

import "fmt"

func main() {
	var x string = "Isto é uma variavel do tipo string"
	nome := "Welliton"

	var (
		teste1 = "welliton"
		teste2 = "mayara"
	)

	casa1, casa2 := "home", "house"

	const constante = "Não muda"

	fmt.Println(constante)
	fmt.Println(casa1)
	fmt.Println(casa2)
	fmt.Println(teste1)
	fmt.Println(teste2)
	fmt.Println(x)
	fmt.Println(nome)
}
