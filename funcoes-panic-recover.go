package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado com sucesso.")
	}
}

func media(x, y int) bool {
	defer recuperar()
	media := (x + y) / 2

	if media < 6 {
		return false
	} else if media > 6 {
		return true
	}

	panic("Média igual a 6 panico")
}

func main() {
	fmt.Println(media(6, 6))
}
