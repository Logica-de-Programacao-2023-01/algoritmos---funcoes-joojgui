package main

// Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.
func media(slice []int) float64 {
	var soma int

	for _, valor := range slice {

		soma += valor

	}
	return float64(soma) / float64(len(slice))
}
