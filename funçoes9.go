package main

import (
	"errors"
	"strings"
)

//Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string.
//Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro.

func Totalpalavras(str string) ([]string, error) {

	if len(str) == 0 {

		return nil, errors.New("A slice está vazia ")

	}

	novastring := strings.Split(str, "")

	return novastring, nil

}
