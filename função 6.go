package main

import (
	"errors"
	"strings"
)

//Escreva uma função que receba um slice de strings como parâmetro e
//retorne uma string com todas as strings concatenadas e separadas por vírgulas.
//Caso o slice esteja vazio, retorne um erro.

func ConcatenaçaoEsep(slice []string) (string, error) {

	if len(slice) == 0 {

		return "", errors.New("A slice é vazia")

	}

	return strings.Join(slice, ","), nil

}
