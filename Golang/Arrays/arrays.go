package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	//da para deixar assim onde ele nao tem um valor fixo, porem precisa informar os valores
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	//o slice no Go acaba sendo mais usado por ser mais flexivel com a qunatidade de itens no arrays
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	//o append insere novos valores no array
	slice = append(slice, 16)
	fmt.Println(slice)

	//o slice consegue pegar os valores em intervalos dos arrays
	slice2 := array2[1:3]
	fmt.Println(slice2)

	//se mudar no array muda no slice
	array2[1] = 4
	fmt.Println(slice2)

	//Arrays internos
	//10 = tamanho / 15 = capacidade maxima
	//função make cria um array com os dados que informamos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length(tamanho)
	fmt.Println(cap(slice3)) //capacidade

	//sempre que passar da capacidade ele ira dobrar e criar mais espaços
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)

	fmt.Println(len(slice3)) // length(tamanho)
	fmt.Println(cap(slice3)) //capacidade

	//se nao por a capacidade maxima ele usa o valor da capacidade
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length(tamanho)
	fmt.Println(cap(slice4)) //capacidade

}
