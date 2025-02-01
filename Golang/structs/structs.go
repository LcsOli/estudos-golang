package main

import "fmt"

func main() {
	fmt.Println("Arquivos Structs")

	var u usuario
	u.nome = "Lucas"
	u.idade = 28
	fmt.Println(u)

	//Voce pode inserir dados de um struct e inserir em outro struct o que ira gerar um dentro do outro {{}}
	endExemplo := endereco{"Rua 1", 0}
	//Outra forma de fazer a criação é assim:
	usuario2 := usuario{"Lucas", 28, endExemplo}
	fmt.Println(usuario2)

	//Se caso precisar pode utilizar deste jeito quando não tem a outra informação
	usuario3 := usuario{nome: "Lucas"}
	fmt.Println(usuario3)
}

//Lembra a orientação a objetos
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}
