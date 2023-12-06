package main

import "fmt"

func somar(x int, y int) int {
	return x + y
}

func calcular(x int, y string) (int, string) {
	return x + 5, y + " ok"
}

func main() {
	result := somar(5, 10)
	fmt.Println("soma: ", result)
	var f = func(txt string) string {
		return txt
	}
	resultado := f("texto")
	fmt.Println(resultado)

	resultado1, resulado2 := calcular(5, "welliton")
	fmt.Println(resultado1, resulado2)
}
