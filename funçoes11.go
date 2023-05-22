package main

import "errors"

//Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número primo
//e falso caso contrário. Caso o número seja negativo, retorne um erro

func Primos(num int) (bool, error) {

	if num < 0 {

		return false, errors.New("Error")

	}

	if num/num == 0 {

	}

}
