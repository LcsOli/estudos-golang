package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2 "
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "var 3"
		variavel4 string = "var 4 "
	)
	fmt.Println(variavel3, variavel4)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)
}
