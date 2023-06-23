package main

import (
	"errors"
	"fmt"
)

func main() {
	strSlice := []string{"apple", "banana", "cherry", "orange", "grape"}

	filteredSlice, err := filterStrings(strSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings com mais de 5 caracteres:", filteredSlice)
	}
}

func filterStrings(strSlice []string) ([]string, error) {
	if len(strSlice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	filteredSlice := make([]string, 0)

	for _, str := range strSlice {
		if len(str) > 5 {
			filteredSlice = append(filteredSlice, str)
		}
	}

	return filteredSlice, nil
}
