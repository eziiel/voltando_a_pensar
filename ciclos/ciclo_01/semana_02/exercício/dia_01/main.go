package main

import (
	"fmt"
)

func main() {
	conexoes := map[string][]string{
    "Ana":      {"Bruno", "Carlos"},
    "Bruno":    {"Ana", "Daniel", "Eduardo"},
    "Carlos":   {"Ana", "Fernanda"},
    "Daniel":   {"Bruno"},
    "Eduardo":  {"Bruno"},
    "Fernanda": {"Carlos"},
	}

	fila := []string{"Ana"}
	comunicados := map[string]bool{}

		// pessoa := fila[0]
		// fila = fila[1:]

		// comunicados[pessoa] = true
		// fila = append(fila, conexoes[pessoa]...)

		// if len(comunicados) > 0 {
		// 	pessoa2 := fila[0]
		// 	fila = fila[1:]
		// 	comunicados[pessoa2] = true
		// 	fila = append(fila, conexoes[pessoa2]...)
		// }

		// if len(comunicados) > 0 {
		// 	pessoa3 := fila[0]
		// 	fila = fila[1:]
		// 	comunicados[pessoa3] = true
		// 	fila = append(fila, conexoes[pessoa3]...)
		// }

		// fmt.Println("Fila:", fila, "Comunicados:", comunicados)

	for len(fila) > 0 {
				pessoa := fila[0]
				fila = fila[1:]

				if comunicados[pessoa] {
					continue
				} 
				
				comunicados[pessoa] = true
				fila = append(fila, conexoes[pessoa]...)
			}
			fmt.Println("Fila:", fila, "Comunicados:", comunicados)
}