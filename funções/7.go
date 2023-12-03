package main

import "fmt"

//Crie uma função que receba um slice de inteiros e uma função como parâmetros.
//A função deve aplicar a função recebida em cada elemento do slice e retornar um
//novo slice com os resultados. Caso o slice esteja vazio, retorne um erro

func main() {
	x := mediaAritmetica([]int{1.0, 4.0, 7.0, 10.0})
	fmt.Println(x)
}

func mediaAritmetica(slice []int) float64 {
	soma := 0.0
	for _, valor := range slice {
		soma += float64(valor)
	}
	media := soma / float64(len(slice))
	return media
}
