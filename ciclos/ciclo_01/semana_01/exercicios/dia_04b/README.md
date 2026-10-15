# Exercício 03 — Parênteses

## Problema

```
Você recebe uma lista de palavras e precisa agrupar as palavras que são anagramas entre si.

Exemplo:

Entrada:
["amor", "roma", "carro", "ramo", "arroz", "omar"]
```

Saída esperada conceitualmente:

```
[
  ["amor", "roma", "ramo", "omar"],
  ["carro"],
  ["arroz"]
]
```

amor, roma, ramo e omar possuem as mesmas letras, apenas em ordens diferentes.

Outro exemplo:

```
["abc", "bca", "teste", "cab", "ettse", "xyz"]

```
→

```
[
  ["abc", "bca", "cab"],
  ["teste", "ettse"],
  ["xyz"]
]
```
A ordem dos grupos não importa.

## O que preciso descobrir
 - O que é um anagrama
 - como mapear as letras? mesmo que eu ache uma palavra com o mesmo length não quer dizer que é um anagrama, mas já é um comparador.
 - letra por letra, loops

...

## Estratégia
 - em primeiro momento nenhuma, fui estudar estrutura de dados.
 - novamente pensando em loop dentro de loop, acho que preciso estudar mais estruturas de dados
 - imagino que vou ter que comparar as palavras, para ver se elas possuem as letras de uma outra palavra
 - talvez ir removendo a letra da palavra em comparação
 
 ```
    roma - amor
    [r, o, m, a] - [a, m, o, r]
    [r] - [a, m, o]
    [o] - [a, m]
    [m] - [a]
    [a] - []

 ```
  - e se eu ordenar as letras em ordem alfábética e um loop ordenado ir removendo?
  - preciso renmover?

  crio uma estrutura chave valor de cara palavra em comparação e vejo se item e a quantidade desses são iguais?

  amor - {a: 1, m: 1, o: 1, r: 1}
  roma - {r: 1, o: 1, m: 1, a: 1} - são iguais!

  carro - {c: 1, a: 1, r: 2}
  ramo - {r: 1, a: 1, m: 1, o: 1} - não são iguais!

    considerando o retorno pedido, apende do item em validação em variáveis sendo ou não anagrama, 
...


## Estado que preciso manter


...

## Estruturas de dados que parecem úteis

...

## Complexidade que imagino

...

## comentário

acho que preciso estudar estruturas de dados
