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
	"fmt"
	"strings"
)

// func main() {
//     fmt.Print("Digite sua mensagem: ")

// 	var palavra string
// 	fmt.Scanln(&palavra)

// 	fmt.Println(palavra)

// 	var letras = strings.Split(palavra, "")

// 	var total int = 1

// 	for i := range letras {
// 		for j := range letras {
// 			if letras[j] == letras[i] {
// 				total++
// 			}
// 		}
// 		if(total == 2 ) {
// 			fmt.Println("a letra: ", letras[i], " é a primeira que aparece somente uma vez!")
// 			return
// 		}
// 		total = 1
// 	}

// 	fmt.Println("Não existe itens que se repetem apenas uma vez!")
// }

func main () {

	var palavra string
	fmt.Scanln(&palavra)

	var letterCount = make(map[string]int)
	var letters = strings.Split(palavra, "")

	fmt.Println("primeira contagem: ", letterCount)

	for i := range letters {
		letterCount[letters[i]] = letterCount[letters[i]] + 1
	}

	fmt.Println("segunda contagem: ", letterCount)
	
	//levando em consideração que é da esquerda para direita
	for i := range letters {
		if letterCount[letters[i]] == 1 {
			fmt.Println("a letra: ", letters[i], " é a primeira que aparece somente uma vez!")
			return
		}	
	}

	fmt.Println("Não existe itens que se repetem apenas uma vez!")
}

