package main

import (
	"fmt"
	"log"
	"meumoduloteste/app"
	"os"
)

func main() {
	fmt.Println("Aplicação sobre a linha de comando")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
