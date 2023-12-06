package main

import "fmt"

func main() {
	fmt.Println("maps")

	maps := map[string]string{
		"nome":       "Welliton",
		"sobre_nome": "Cunha",
	}

	fmt.Println(maps["nome"])

	maps2 := map[string]map[string]string{
		"usuarios": {
			"nome":        "Welliton",
			"ultimo_nome": "Cunha",
		},
		"curso": {
			"nome":   "Facimp",
			"campus": "camus 1",
		},
	}
	fmt.Println(maps2["usuarios"]["ultimo_nome"])
	delete(maps2, "nome")
	maps2["test"] = map[string]string{
		"test": "test",
	}
	fmt.Println(maps2)
}
