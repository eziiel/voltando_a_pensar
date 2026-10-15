package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	palavras := []string{
    "123",
    "321",
    "456",
    "231",
    "654",
    "111",
    "11",
}

  
	grupos := make(map[string][]string)


	for _, palavra := range palavras {
		fmt.Println("palavra: ", palavra)

		letras := strings.Split(palavra, "")
		fmt.Println("letras1: ", letras)

		sort.Strings(letras)
		fmt.Println("letras: ", letras)

		chave := strings.Join(letras, "")
		fmt.Println("chave: ", chave)

		grupos[chave] = append(grupos[chave], palavra)
	}

	fmt.Println(grupos)
}