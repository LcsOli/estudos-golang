package main

import "fmt"

func main() {
	numero := 1

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//voce pode criar uma variavel e ja inserir uma validação
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Numero é maior que 0")
	} else if numero < -10 {
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	//Fora da validação ela não existe mais
	/*fmt.Println(numero2)*/
}
