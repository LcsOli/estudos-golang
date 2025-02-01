package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	// so ira mexer no valor informado, pois como a variavel2 esta referenciando o valor inicial da variavel1 ela permanece igual, mas podemos usar variavel2++ para acrescentar apenas na 2
	variavel1++
	fmt.Println(variavel1, variavel2)

	//ponteiros são referencias de memoria
	var variavel3 int = 100
	//o * mostra que é um ponteiro
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	//ao usar o & mostramos ao sistema que o ponteiro ira pegar o endereço de memoria da variavel
	variavel3 = 0
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	fmt.Println(variavel3, *ponteiro) //desreferenciação

	//se houver alteração o ponteiro ira pegar o mesmo endereço de memoria, porem ao desreferenciar o valor altera como no exemplo abaixo
	variavel3 = 15
	fmt.Println(variavel3, *ponteiro)

}
