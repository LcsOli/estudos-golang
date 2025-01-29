package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Elias gay")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("teste@email.com")
	fmt.Println(erro)
}
