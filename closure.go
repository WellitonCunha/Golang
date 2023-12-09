package main

import "fmt"

func closure() func() {
	texto := "Texto da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "texto da função main"
	fmt.Println(texto)
	novafuncao := closure()
	novafuncao()
}
