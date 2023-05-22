package main

import (
	"errors"
)

// Escreva uma função que receba um slice de inteiros como parâmetro
//e retorne um novo slice com apenas os números pares contidos no slice. Caso o slice esteja vazio, retorne um erro.

func Slicepares(slice []int) ([]int, error) {

	if len(slice) == 0 {

		return nil, errors.New("A slice esta vazia")
	}

	var newslice []int

	for _, slc := range slice {

		if slc%2 == 0 {

			newslice = append(newslice, slc)

		}

	}

	return newslice, nil
}
