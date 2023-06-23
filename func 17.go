package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func main() {
	strSlice := []string{"banana", "abacaxi", "laranja"}

	sortedStr, err := sortStrings(strSlice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings em ordem alfabética:", sortedStr)
	}
}

func sortStrings(strSlice []string) (string, error) {
	if len(strSlice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(strSlice)

	// Junta as strings em uma única string separadas por espaço
	sortedStr := strings.Join(strSlice, " ")

	return sortedStr, nil
}
