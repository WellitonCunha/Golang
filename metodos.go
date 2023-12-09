package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os o %s no banco de dados", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Welliton", 31}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Mayara", 31}
	fmt.Println(usuario2)
	fmt.Println(usuario1.maiorIdade())
}
