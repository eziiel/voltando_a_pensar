package main

import "fmt"


func main() {
	
	// levando em consideração que a ordem do array é a ordem de entrada
	
	
	// for i, numero := range numeros {
		
	// 	// pilha = append(pilha, numeros[len(numeros) - i - 1])
	// 	pilha = append(pilha, numero)
	
	// 	fmt.Println(i, numero)
	// }
	
	
	// for i := len(numeros) - 1; i >= 0; i-- {
	// 		pilha = append(pilha, numeros[i])
	// 	}
	
	// slices.Reverse(numeros)
	numeros := []int{10, 20, 30, 40, 50}
	var pilha []int
	
	// var recursive func() []int
	
	// recursive = func() []int {
	// 	if len(numeros) != 0 {

	// 		pilha = append(pilha, numeros[len(numeros)-1])
		
	// 		numeros = numeros[0:len(numeros)-1]


	// 		recursive()
	// 	}
		
	// 	return []int{}
	// }
	
	// recursive()


	// for i := range numeros {
	// 	// fmt.Println("numero: ", numeros[len(numeros) - 1])
	// 	fmt.Println(i)
	// 	// pilha = append(pilha, numeros[len(numeros)-numero])
		
	// 	numeros = numeros[0:len(numeros) - 1]
	// 	fmt.Println("numeros: ",numeros)
	// }
	
for len(numeros) > 0 {

    // 1. pegar o elemento do topo
    var numero = numeros[len(numeros)-1]
		pilha = append(pilha, numero)
    // 2. fazer alguma coisa com ele
    // 3. remover o topo
    numeros = numeros[:len(numeros)-1]
}

	fmt.Println("pilha: ", pilha)
	fmt.Println("numeros: ", numeros)

}

