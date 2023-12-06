package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("ola")

	p1 := pessoa{"Welliton", "Cunha", 31, 1.70}
	p2 := estudante{p1, "Sistema de Informação", "facimp"}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.nome)
}
