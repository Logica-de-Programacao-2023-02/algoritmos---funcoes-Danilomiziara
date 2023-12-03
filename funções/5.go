package main

import "fmt"

//Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do
//primeiro elemento igual ao valor no slice. Caso não encontre, retorne -1.

func main() {
	numeros := []int{3, 7, 5, 2, 9}
	posicao := EncontrarPosicaoElemento(numeros, 5)
	fmt.Println(posicao)
}

func EncontrarPosicaoElemento(slice []int, valor int) int {
	for i, v := range slice {
		if v == valor {
			return i
		}
	}

	return -1
}
