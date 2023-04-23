package main

//Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.

func Segundomaiorvalor(slice []int) int {

	maior := slice[0]
	segundoMaior := slice[1]

	for _, num := range slice {
		if num > maior {
			segundoMaior = maior
			maior = num
		} else if num > segundoMaior && num != maior {
			segundoMaior = num
		}
	}

	return segundoMaior

}
