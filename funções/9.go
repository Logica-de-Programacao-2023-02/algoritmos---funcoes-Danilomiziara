//Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as
//palavras contidas na string. Considere que as palavras são separadas por espaços em branco.
//Caso a string seja vazia, retorne um erro.

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	frase := "que coisa incrivel"
	palavras, err := RetornarPalavras(frase)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(palavras)
}

func RetornarPalavras(frase string) ([]string, error) {
	if frase == "" {
		return nil, errors.New("A String está vazia")
	}

	palavras := strings.Split(frase, " ")
	return palavras, nil
}
