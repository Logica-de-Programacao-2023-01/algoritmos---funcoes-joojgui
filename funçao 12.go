package main

import (
	"errors"
	"strings"
	"unicode"
)

//Escreva uma função que receba um slice de strings como parâmetro e
//retorne uma string contendo apenas as strings que começam com uma letra maiúscula.
//Caso o slice esteja vazio, retorne um erro.

func Prefixomaisuculos(slice []string) (string, error) {

	if len(slice) == 0 {

		return "", errors.New("O slice esta vazio :")

	}

	var slicetotal []string

	for _, str := range slice {
		firstChar := rune(str[0])
		if unicode.IsUpper(firstChar) {
			slicetotal = append(slicetotal, str)
		}
	}

	return strings.Join(slicetotal, " "), nil
}
