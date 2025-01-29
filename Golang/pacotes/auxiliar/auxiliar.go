package auxiliar

import "fmt"


// escrever nome da função com letra minuscula apenas os arquivos com o mesmo package vão poder utilizar, caso esteja em maiscula todos podem usar a função //

//OBS: sempre que tiver algo importado ele precisa ter um comentario //
func Escrever()  {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}