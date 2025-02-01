package main

import "fmt"

func main() {
	dia := diasemana(1)
	fmt.Println(dia)
}

func diasemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Digite um valor entre 1 e 7"
	}
}

//pode ser feito assim tambem
/*func diasemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	}
}*/

//vai para a proxima condição
/*	switch numero {
	case 1:
		return "Domingo"
		fallthrough*/
