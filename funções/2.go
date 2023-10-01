package main

import (
	"fmt"
	"strings"
)

// Crie uma função que receba uma string e retorne a quantidade de vogais presentes.
func main() {
	texto := "gabriel EU te amo"
	qntvogais := qv(texto)
	fmt.Println(qntvogais)

}
func qv(texto string) int {
	texto = strings.ToUpper(texto)
	vogais := "AEIOU"
	contador := 0
	for _, char := range texto {
		if strings.ContainsAny(string(char), vogais) {
			contador++
		}
	}
	return contador
}
