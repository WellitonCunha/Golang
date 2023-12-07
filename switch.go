package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Dmingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia Inválido."
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Dmingo"
		fallthrough // ele obriga a condição pular para o proximo case direto no return
		// nesse caso especifico não vai pular porque o return encerra aplicação mas se
		// fosse variavel passava tranquilamente
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sábado"
	default:
		return "Dia Inválido."
	}
}

func main() {
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana2(6))
}
