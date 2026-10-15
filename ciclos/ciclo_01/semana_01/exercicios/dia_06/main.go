// aquecimanto

package main

import (
	"fmt"
	"math/rand"
)


func main() {
	fila := []string{
		"Bruno",
		"Carlos",
		"Cristiana",
		"Daniel",
	}

	aleatorias := []string{	
		"Jose",
		"Pedro",
		"Joana",
		"Carimbo",
		"Enzo",
		"Joanesburgo",
		"Arnaldo",
		"Pimpolho",
		"Raimundo",
	}


	for len(fila) > 0 {
		
		cliente := fila[0]

		fila = fila[1:]
		
		if len(aleatorias) > 0 {
			addClient := rand.Intn(len(aleatorias))
			if rand.Intn(100) < 50 {
				fila = append(fila, aleatorias[addClient])
				aleatorias = append(aleatorias[:addClient], aleatorias[addClient+1:]...)
			}
		}


		fmt.Println("cliente em atendimento", cliente)
		fmt.Println("clientes restantes", fila)
	}
}

