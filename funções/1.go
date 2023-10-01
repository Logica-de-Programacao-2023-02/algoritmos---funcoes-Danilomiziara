package main

import "fmt"

// Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.
func main() {
	numero := []int{2, 10}
	media := ma(numero)
	fmt.Println(media)

}
func ma(slice []int) float64 {
	if len(slice) == 0 {
		return 0.0
	}
	soma := 0
	for _, i := range slice {
		soma += i
	}
	return float64(soma) / float64(len(slice))
}
