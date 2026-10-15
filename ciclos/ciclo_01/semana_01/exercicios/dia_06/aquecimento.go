
package main

import "fmt"


func aquecimento() {
	clientes := []string{
    "Ana",
    "Bruno",
    "Carlos",
    "Daniel",
}


	for len(clientes) > 0 {
		cliente := clientes[0]
		clientes = clientes[1:]
		fmt.Println("cliente em atendimento", cliente)
		fmt.Println("clientes restantes", clientes)
	}
}

