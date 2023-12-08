package main

import "fmt"

func main() {

	result := func(texto string) string {
		return fmt.Sprintf("Olá,  %s ", texto)
	}("Welliton")

	fmt.Println(result)
}
