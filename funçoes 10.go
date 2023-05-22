package main

import (
	"errors"
	"sort"
)

//Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os valores ordenados
//de forma crescente. Caso o slice esteja vazio, retorne um erro.

func Formacrescente(slice []int) ([]int, error) {

	if len(slice) == 0 {

		return nil, errors.New("A slice esta vazia")

	}

	sortedSlice := make([]int, len(slice))
	copy(sortedSlice, slice)
	sort.Ints(sortedSlice)

	return sortedSlice, nil

}


