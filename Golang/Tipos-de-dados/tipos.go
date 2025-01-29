package main

import (
	"errors"
	"fmt"
)

// tipos de int = int(pega o valor do seu sistema ex: x32 ou x64), int8, int16, int32 e int64
func main() {
	var valor int = 100
	fmt.Println(valor)

	//diferença de uint para int é que ele não deixa valores negativos
	var valor2 uint32 = 32
	fmt.Println(valor2)

	//alias
	// int32 = rune
	var valor3 rune = 123456
	fmt.Println(valor3)

	//uint8 = byte
	var valor4 byte = 123
	fmt.Println(valor4)

	//float32, float64
	var numero1 float32 = 123.45
	fmt.Println(numero1)

	var numero2 float64 = 12313132131.44
	fmt.Println(numero2)

	var str string = "Texto"
	fmt.Println(str)

	//pega o numero da tabela ascii
	char := 'A'
	fmt.Println(char)

	//fica vazio pois não tem nada salvo
	var texto string
	fmt.Println(texto)

	// se nao especificar o valor sempre sera falso
	var booleano1 bool
	fmt.Println(booleano1)

	//informa o erro que aconteceu / errors é o pacote
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
