package main

import (
	"fmt"
	"slices"
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

	var montes = make(map[string][]string)

	for _, palavra := range palavras {

		var separarPalavra = strings.Split(palavra, "")

		var ordenarPalavra = separarPalavra

		slices.Sort(ordenarPalavra)

		var juntar = strings.Join(ordenarPalavra, "")
		
		
		montes[juntar] = append(montes[juntar], palavra)
	}

	var resultadoNãoTem [][]string
	var resultadoTem [][]string

	for _, palavra := range montes {



		if len(palavra) > 1 {
			resultadoTem = append(resultadoTem, palavra)
		}else{
			resultadoNãoTem = append(resultadoNãoTem, palavra)
		}
	}

	fmt.Println("resultadoTem: ", resultadoTem)
	fmt.Println("resultadoNãoTem: ", resultadoNãoTem)
}