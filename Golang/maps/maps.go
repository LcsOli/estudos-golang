package main

import "fmt"

func main() {

	// o valor interno é o valor que vai ser as chaves que estão nele e o de fora referenciando o tipo dos valores
	usuario := map[string]string{
		"nome":      "Lucas",
		"Sobrenome": "Oliveira",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Lucas",
			"ultimo":   "Oliveira",
		},
		"curso": {
			"nome":   "TI",
			"campus": "Unip",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
