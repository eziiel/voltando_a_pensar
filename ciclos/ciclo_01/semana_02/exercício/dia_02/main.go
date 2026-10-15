package main

import (
	"fmt"
)

type PessoaFila struct {
	nome			string
	distancia	int
}

func main(){
	conexoes := map[string][]string{
    "Ana":      {"Bruno", "Carlos"},
    "Bruno":    {"Ana", "Daniel", "Eduardo"},
    "Carlos":   {"Ana", "Fernanda"},
    "Daniel":   {"Bruno"},
    "Eduardo":  {"Bruno"},
    "Fernanda": {"Carlos"},
	}

	var pessoaInicio string
	fmt.Scanln(&pessoaInicio)
	
	var pessoaFim string
	fmt.Println(calculaDistancia(conexoes, pessoaInicio, pessoaFim))
}

func calculaDistancia(conexoes map[string][]string, pessoaInicio string, pessoaFim string) string {
	fila := []PessoaFila{
		{pessoaInicio, 0},
	}

	comunicados := map[string]bool{}
	
	for len(fila) > 0 {
		pessoa := fila[0]
		fila = fila[1:]
		
			if pessoa.nome == pessoaFim {
				return fmt.Sprintf("%s está a %d passos de distância", pessoa.nome, pessoa.distancia)
			}

			if comunicados[pessoa.nome] {
					continue
			} 

			comunicados[pessoa.nome] = true

			for _, pessoaConectada := range conexoes[pessoa.nome] {
				fila = append(fila, PessoaFila{
					nome:      pessoaConectada,
					distancia: pessoa.distancia + 1,
				})
			}
		}
		return  "Não encontrado"
}