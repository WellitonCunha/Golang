package main

import "fmt"

var n int

func init() {
	fmt.Println("Inicia Primeira que a função principal")
	n = 10
}

func main() {
	fmt.Println("Função principal")
	fmt.Println(n)
}
