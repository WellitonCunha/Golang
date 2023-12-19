package main

import (
	"fmt"
	"introducao-teste/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Rua Paulista")
	fmt.Println(tipoDeEndereco)
}
