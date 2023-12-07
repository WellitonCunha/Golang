package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0
	for i < 10 {
		i++
		time.Sleep(time.Second)
		fmt.Println("Inclementando...", i)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("Inclementando...", j)
	}

	nomes := [3]string{"Welliton", "Mayara", "Delphino"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, nome := range "Welliton" {
		fmt.Println(indice, string(nome))
	}

	usuario := map[string]string{
		"nome":  "Welliton",
		"Idade": "31",
	}

	for indice, nome := range usuario {
		fmt.Println(indice, nome)
	}

	for {
		fmt.Println("Executando para sempre")
	}
}
