package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexador(escrever("Olá mundo1"), escrever("Olá mundo2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexador(canalDeEentrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEentrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
