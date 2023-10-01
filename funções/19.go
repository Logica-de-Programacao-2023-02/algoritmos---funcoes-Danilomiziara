package main

// Crie uma função que receba um número inteiro como parâmetro e retorne um novo slice contendo
// todos os números primos menores ou iguais a ele. Caso o número seja negativo, retorne um erro.
func main() {

}
func primos(num int) bool {
	if num <= 1 {
		return false
	} else if num <= 3 {
		return true

	} else if num%2 == 0 {
		return false

	}
	i := 6
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
}
