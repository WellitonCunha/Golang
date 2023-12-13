package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá maundo.")
	escrever("Segundo Print")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
