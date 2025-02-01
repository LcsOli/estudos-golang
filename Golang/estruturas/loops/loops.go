package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		//somente usando para mostrar cada ação por segundo
		time.Sleep(time.Second)
		i++
		fmt.Println("Incrementando i", i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j", j)
		time.Sleep(time.Second)
	}

	fmt.Println("-----------------------------------------")

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	//se nao informar o que quer quando fizer um range dentro de uma string ele ira pegar os valores correspondentes da tabela ASCII
	for indice, letra := range "Letras" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Silva",
	}
	//podemos usar os dados de um map para usar no for porem nao podemos usar dentro de um struct
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for true {
		fmt.Println("Executando...")
	}
}
