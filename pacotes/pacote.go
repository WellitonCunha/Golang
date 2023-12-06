package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("teste")
	auxiliar.Auxiliar()
	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
