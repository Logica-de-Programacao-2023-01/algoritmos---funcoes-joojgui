package main

//Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.

func Concatenação(str []string) string {

	var resultado string

	for _, soma := range str {

		resultado += soma

	}

	return resultado

}


