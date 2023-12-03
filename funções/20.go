package main

import (
	"errors"
	"fmt"
	"strings"
)

//Escreva uma função que receba um slice de strings como parâmetro e retorne um novo
//slice contendo apenas as strings que possuem mais de 5 caracteres.
//Caso o slice esteja vazio, retorne um erro.

func main() {
	Frase := "aumento da atividade eletrica do cortex frontal esquerdo"
	cinco, err := SepararPalavras(Frase)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cinco)

}
func SepararPalavras(slice string) ([]string, error) {
	tamanho := strings.Fields(slice)
	cinco := []string{}
	if len(tamanho) == 0 {
		return nil, errors.New("A slice está vazia")
	}
	for _, palavra := range tamanho {
		if len(palavra) >= 5 {
			cinco = append(cinco, palavra)
		}

	}
	return cinco, nil
}
