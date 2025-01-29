package main

import "fmt"

//função é um tipo, então voce consegue fazer varias coisas com ela
//me lembra bem uma classe do java
func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		//retorna o valor
		return txt
	}

	f("Texto da função")

	// pode usar o valor direto assim ou
	fmt.Println(calculosMatematicos(1, 2))

	//assim, mas voce precisa usar as variaveis informadas senão insira o _ na variavel que não ira utilizar
	//exemplo: resultadosoma, _ := calculosmatematicos(1,1)
	//fmt.Println(resultadosoma) assim ele roda sem dar erro onde não ira precisar informar as duas
	resultadosoma, resultadosubtracao := calculosMatematicos(3, 4)
	fmt.Println(resultadosoma, resultadosubtracao)

}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//podemos utilizar os calculos assim onde tem mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
