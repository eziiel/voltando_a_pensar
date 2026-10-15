# Exercício 06 — Fila (Queue) - Aquecimento

Você tem:
```
clientes := []string{
    "Ana",
    "Bruno",
    "Carlos",
    "Daniel",
}
```

Eles chegaram nessa ordem.

Simule um atendimento onde o sistema sempre atende quem está há mais tempo esperando.

A saída deve ser:

```
Atendendo Ana
Atendendo Bruno
Atendendo Carlos
Atendendo Daniel
```

Mas existe uma regra: a cada atendimento, a pessoa precisa sair da fila.

Então queremos observar algo parecido com:

fila inicial:
```
[Ana Bruno Carlos Daniel]
```

```
atende Ana
[Bruno Carlos Daniel]
```

```
atende Bruno
[Carlos Daniel]
```

```
atende Carlos
[Daniel]
```

```
atende Daniel
[]
```


- Não use slices. Reverse, não crie outro slice com a ordem pronta e não precisa pesquisar “como implementar queue em Go”.

- Você já sabe manipular slices o suficiente para descobrir.

A pergunta que inicia o exercício é:


## Problema
Na stack você removia pilha[len(pilha)-1]. Se numa fila precisamos remover "Ana" de ["Ana", "Bruno", "Carlos", "Daniel"], qual posição do slice representa o próximo a sair?


## O que preciso descobrir
Qual posição do slice representa o próximo a sair?


## Estratégia
- [x] slice 
- [x] Reverse
- [ ] loop while


## Estado que preciso manter
- item atual


## Estruturas de dados que parecem úteis
- Queue (fila)


## Complexidade que imagino
- On


## adicional
 Em uma fila é comum atender o primeiro, ela sai da fila, e, enquanto isso o que resta da fila ainda tem uma sequencia, agora o ponto é: e se aparecer alguem na fila?


## Comentários
  Preciso de alguma forma adicionar uma pessoa, em primeiro momento quero:
  
  - tem que ser no final (último da fila)
  - tem que ser de uma forma controlada para não gerar infinit loop

  - (alter) queria que fosse não uma adicão controlada, para isso criar uma lista de pessoas, e usar um random para pegar uma aleatória, pode ou não ser todas, mas isso me faz pensar em que momento vou adicionar, a cada pessoa? tem alguma outra ordem aleatoria para add?
  