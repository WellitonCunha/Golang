package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"Welliton", "Mayara", "Mãe", "Pai", "Irmão"}
	fmt.Println(array2)

	// com os três ponto o array é criado de acordo a quantidade de dados que foi inserido
	array3 := [...]string{"Welliton", "Mayara", "Mãe", "Pai", "Irmão"}
	fmt.Println(array3)

	slices := []int{10, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slices)

	fmt.Println(reflect.TypeOf(slices))
	fmt.Println(reflect.TypeOf(array3))

	slices = append(slices, 1)
	fmt.Println(slices)

	slices2 := array2[1:3]
	fmt.Println(slices2)

	array2[1] = "Posição alterada"
	fmt.Println(slices2)

	// arrays internos
	fmt.Println("--------")
	slices3 := make([]float32, 10, 11)
	fmt.Println(slices3)

	slices3 = append(slices3, 5)
	slices3 = append(slices3, 6)

	fmt.Println(len(slices3)) // tamanho
	fmt.Println(cap(slices3)) // capacidade

	// resumindo a diferença de arrays x slices:
	// arrays tem tamanho fixo
	// slices tem tamanho dinamico
}
