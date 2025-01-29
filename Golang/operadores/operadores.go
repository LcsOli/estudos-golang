package main

import "fmt"

func main() {

	//ARITMETICOS
	// +,-,/,*,%(RESTO DA DIVISAO)

	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 2
	mult := 7 * 2
	resto := 5 % 2

	fmt.Println(soma, sub, div, mult, resto)

	//ATRIBUICAO
	//Lembrando que se criar uma variavel com := não consegue reutilizar a variavel ela fica como se fosse a const ou final em java
	var variavel string = "String"
	variavel2 := "String2"
	fmt.Print(variavel, variavel2)

	//RELACIONAIS
	//SEMPRE RETORNAM BOOLEANO
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2) //valida se é o mesmo tipo
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2) //diferente

	//OPERADORES LOGICOS
	fmt.Println("------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNITARIOS
	numero := 10
	numero++     //acrescenta 1
	numero += 15 //acrescenta o valor apos igual (10 + 15)
	fmt.Println(numero)

	numero--     //deduz 1
	numero -= 20 // deduz pelo valor
	numero *= 3  // multiplica pelo valor informado
	numero /= 5  // divide pelo valor informado
	numero %= 3  //calcula o resto da operação
	fmt.Println(numero)

}
