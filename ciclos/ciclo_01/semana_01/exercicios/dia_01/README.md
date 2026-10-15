# Exercício 01 — Frequência

## Problema

Implemente uma função que recebe uma string e retorna o caractere que mais aparece nela.

## Minha interpretação

Vou fazer um loop no length da palava ex: casa, 4, na primeira iteração do loop i = 0 = c
depois disso vou salvar a palavra[i], agora vou fazer um novo loop no range de palavra
item1 = 0
item2 = 0
se o letras[j] for igual a letra[i] somo mais um ao item
condicional item1 > item2? se sim item2 = item1

depois printo item2


## Estratégia

abrir uma funcao e pedir entrada de dados, ir printando o que ta aparecendo, caso tenha uma ideia e 
não souber como fazer vou pesquisar

ex: como splita uma var em go?
...

## Estrutura de dados

var
array
loops
iteracão
string
caracteres
funcoes
números

...

## Complexidade esperada
O(n^2) (auto complete droga :/)
...

## Casos que preciso testar

- string vazia
- espaços
- números
- letras
- símbolos

ps. pesquisei como faz checkbox no .md, auto complete é um porre as vez pq completa, onde já se viu kkkkkk
...

## comentários de uso
se for uma letra/caracter que já apareceu, então necessáriamente preciso contabiliazar, mas.... a === à?
letra então pode ser A-aZ-z

vamos com calma

...

colocar a letra na primeira iteração deixa um pouco confuso, mas eu colocaria tanto a letra validade quando os contadores fora das iterações?

kaskaskasksakk: 
…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
escreva uma palavra
casa
a 6

... 

acho que iteração não da certo 'pqp, to muito enferrujado' ou não sei programar...

...

 foi mas foi do pior jeito possivel :/

```
…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
escreva uma palavra
            
nenhuma letra encontrada

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
escreva uma palavra
                              a a 
Todas as letras aparecem uma vez 1

…/dev/treinamento ❯ a 
Choose default agent with: omarchy default agent <name>

…/dev/treinamento ✗ a 
Choose default agent with: omarchy default agent <name>

…/dev/treinamento ✗ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
escreva uma palavra
                                       211212212121121212222
o caracter:  2  aparece  12  vezes

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
escreva uma palavra
                      @##$$$$$$$$$$$$$$$$$$$$$$$$$   asd
o caracter:  $  aparece  25  vezes

…/dev/treinamento ❯   asd
bash: command not found: asd

…/dev/treinamento ✗ $$$$$$$$$$$$$$$$$$$$$$$$$
bash: command not found: 928803928803928803928803928803928803928803928803928803928803928803928803$

```

### comentários finais

fiquei meio perdido, acho que por não saber tanto de go, ou pela lógica, prints saiam em lugares desnecessarios, mas ajudou muito a me localizar, esqueci que tinha contador e precisava zerar sempre ksksak, tirar variáveis das iterações ajudou. mantive a ideia de dupla iteração. mas não sei se to satisfeito, demorei :/