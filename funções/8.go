package main

import "fmt"

// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com apenas
// os números pares contidos no slice.
// Caso o slice esteja vazio, retorne um erro.

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares, err := ApenasPares(numeros)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(pares)
}

func ApenasPares(slice []int) ([]int, error) {
	pares := []int{}
	for _, valor := range slice {
		if valor%2 == 0 {
			pares = append(pares, valor)

		}
	}
	return pares, nil
}
