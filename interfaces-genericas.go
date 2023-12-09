package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Welliton")
	generica(1)
	generica(true)
	generica(1.7)
}
