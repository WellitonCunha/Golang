package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func main() {
	fmt.Println("Olá")
	var u usuario
	u.nome = "Welliton"
	u.idade = 21
	fmt.Println(u)

	usuario2 := usuario{"Mayara", 32}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 32}
	fmt.Println(usuario3)
}
