package main

//Crie uma função que receba um slice de inteiros e um valor inteiro,
//e retorne a posição do primeiro elemento igual ao valor no slice. Caso não encontre, retorne -1.

func EncontrarValor(nums []int, valor int) int {
	for i, num := range nums {
		if num == valor {
			return i
		}
	}

	return -1
}
