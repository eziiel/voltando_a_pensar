//
// Exercício 01 — Frequência
// Implemente uma função que recebe uma string e retorna o caractere que mais aparece nela.

// vou fazer um loop no length da palava ex: casa, 4, na primeira iteração do loop i = 0 = c
// depois disso vou salvar a palavra[i], agora vou fazer um novo loop no range de palavra
// item1 = 0
// item2 = 0
// se o letras[j] for igual a letra[i] somo mais um ao item
// condicional item1 > item2? se sim item2 = item1

// depois printo item2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Digite sua mensagem: ")

	fmt.Print("Digite sua mensagem: ")
    if scanner.Scan() {
        palavra := scanner.Text() // aqui você tem o texto digitado
        fmt.Printf("Você digitou: %s\n", palavra)
        fmt.Printf("tamanho: %d\n", len(palavra))
    }

	var palavra string
	fmt.Println("escreva uma palavra")
	fmt.Scanln(&palavra)

	var letraPonteira string
	var totalCount = 0
	var totalFinal = 0

	var letras = strings.Split(palavra, "")

	if len(letras) == 0 {
		fmt.Println("nenhuma letra encontrada")
		return
	}


	for i := range letras {
		
		for j := range letras {
			if letras[j] == letras[i] {
				totalCount = totalCount + 1
			}
			if totalCount > totalFinal {
				totalFinal = totalCount
				letraPonteira = letras[i]
			}
		}
		totalCount = 0
	}
	if totalFinal == 1 {
		fmt.Println("Todas as letras aparecem uma vez", totalFinal)
	} else {
		fmt.Println("o caracter: ", letraPonteira, " aparece ", totalFinal, " vezes")
	}
}