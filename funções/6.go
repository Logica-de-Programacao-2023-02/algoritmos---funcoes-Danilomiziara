// Escreva uma função que receba um slice de strings como parâmetro e
// retorne uma string com todas as strings concatenadas e separadas
// por vírgulas. Caso o slice esteja vazio, retorne um erro.
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	palavras := []string{"Oi", "meu", "amigo"}
	resultado, err := ConcatenarComVirgulas(palavras)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}

func ConcatenarComVirgulas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice está vazio")
	}

	return strings.Join(slice, ","), nil
}
