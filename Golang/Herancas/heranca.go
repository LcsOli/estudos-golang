package main

import "fmt"

//No go não temos a herança então com os structs podemos criar algo mais proximo

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Lucas", 28, 175}
	fmt.Println(p1)

	e1 := estudante{p1, "TI", "UNIP"}
	fmt.Println(e1)

}

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}
