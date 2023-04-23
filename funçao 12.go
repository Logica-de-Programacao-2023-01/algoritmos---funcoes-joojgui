package main

import "errors"

//Escreva uma função que receba um slice de strings como parâmetro e
//retorne uma string contendo apenas as strings que começam com uma letra maiúscula.
//Caso o slice esteja vazio, retorne um erro.

func Prefixomaisuculos(slice []string) (string, error) {

	if len(slice) == 0 {

		return "", errors.New("O slice esta vazio :")

	}

}
