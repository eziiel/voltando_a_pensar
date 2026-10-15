package main

import (
	"fmt"
	"strings"
)



func main() {

	fmt.Println("oi!")

	var palavra string
	fmt.Scanln(&palavra)

	// var parOpen = 0
	// var parClose = 0

	var letters = strings.Split(palavra, "")

	// for i := range letters {

	// 	if letters[i] == "(" {
	// 		parOpen++
	// 	}
	// 	if letters[i] == ")" {
	// 		parClose++
	// 	}
	// }
		
	// 	if parOpen != parClose {
	// 		fmt.Println("a expressão não é valida!2")
	// 		return
	// 	}
	// 	fmt.Println("a expressão é valida!")

	var parValidator []int
	// var parClose []int

	// var letters = strings.Split(palavra, "")

	// for i := range letters {

	// 	if letters[i] == "(" {
	// 		parOpen = append(parOpen, i)
	// 	}

	// 	if letters[i] == ")" {
	// 		parClose = append(parClose, i)
	// 	}
	// }

	// var response []int

	// fmt.Println("parOpen, parClose: ", parOpen, parClose)
	// fmt.Println("len(parOpen), len(parClose): ", len(parOpen), len(parClose))

	// if len(parOpen) < len(parClose) {
	// 	response = parClose[len(parOpen):len(parClose)]
	// 	fmt.Println("nao é valida", response)
	// 	return
	// }
	
	// if len(parOpen) > len(parClose) {	
	// 	response := parOpen[len(parClose):len(parOpen)]
	// 	fmt.Println("nao é valida2", response)
	// 	return
	// }	
	// 	fmt.Println("a expressão é valida!")


	
	for i := range letters {
		if letters[i] == "(" {
			parValidator = append(parValidator, i)
		}
		
		if letters[i] == ")" {
			if len(parValidator) == 0 {
				parValidator = append(parValidator, i)
				fmt.Println("a expressão não é valida! falta abertura! ", parValidator) 
				return
			}
			parValidator = parValidator[:len(parValidator)-1]		
		}
	}

	if len(parValidator) == 0 {
		fmt.Println("a expressão é valida!", parValidator)
		return
	}

	fmt.Println("a expressão não é valida! falta fechamento!", parValidator)
}