package main

//Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os valores
//ordenados de forma crescente. Caso o slice esteja vazio, retorne um erro.

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	numeros := []int{10, 5, 9, 34, 23}
	ordenados, err := NumerosCrescentes(numeros)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ordenados)
}

func NumerosCrescentes(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazio")
	}

	ordenados := make([]int, len(slice))
	copy(ordenados, slice)
	sort.Ints(ordenados)

	return ordenados, nil
}
