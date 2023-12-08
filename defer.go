package main

import "fmt"

func texto1() {
	fmt.Println("defer")
}

func texto2() {
	fmt.Println("Segundo texto")
}

func media(x, y float32) (float64, string) {
	media := (x + y) / 2

	if media <= 6 {
		return float64(media), "Reprovado."
	}

	return float64(media), "Aprovado."
}

func main() {
	// defer retarda a exceção até o ultimo minuto, em outras palavras ele vai ser sempre
	// o último a ser executado
	defer texto1()
	texto2()

	fmt.Println(media(3, 4))
}
