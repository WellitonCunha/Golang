package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8 int16 int32 int64
	var numero int = 10 // int sem o numero de bits vai usar a quantidade de bits computador, tipo 32 bits ou 64 bits
	fmt.Println("ola", numero)

	// uint8 uint16 uint32 uint64
	// uint -> não aceita sinais tipo -1000
	var x uint = 100
	fmt.Println(x)

	//alias
	// int32 = rune
	// uint8 = byte

	var var32 rune = 200
	var bteste byte = 30
	fmt.Println(var32)
	fmt.Println(bteste)

	var numerReal1 float32 = 200.89
	var numerReal2 float64 = 307777777777.65
	fmt.Println(numerReal1)
	fmt.Println(numerReal2)

	var booleano bool = false
	fmt.Println(booleano)

	var erro error = errors.New("Erro de teste")
	fmt.Println(erro)

}
