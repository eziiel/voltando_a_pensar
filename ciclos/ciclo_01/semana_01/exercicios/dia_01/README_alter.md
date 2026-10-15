## Resultado

Consegui implementar uma solução O(n²)
usando dupla iteração.

## Dificuldades

 - Esqueci de resetar o contador entre as iterações
 - tive dificuldades em acompanhar o estado das variáveis
 - tive dúvidas sobre string/caracteres em go
 - demorei mais que o esperado

## O que percebi

Minha maior dificuldade não foi escrever os loops,
mas acompanhar o estado das variáveis durante eles.

Preciso melhorar minha capacidade de visualizar
o que acontece a cada iteração.

## Próxima investigação

- Como Go representa caracteres em strings?
- Diferença entre byte, rune e string.
- Como funciona strings.Split.

Testei:

AaáàãâA
çé😀çé

Caracteres visualmente diferentes são tratados
como diferentes pelo programa.

Também consegui trabalhar com caracteres Unicode
como ç, é e 😀.

## Empates

Exemplo:

aabb

a → 2
b → 2

Pergunta:

O que deve retornar?

Possibilidades:

- primeiro caractere encontrado
- último caractere encontrado
- todos os caracteres empatados
- outro critério

Decisão:
...

## Observação

Ao utilizar as setas durante a entrada, sequências como
^[[A e ^[[D aparecem como parte da entrada.

Provavelmente são sequências de escape do terminal.
Investigar posteriormente.

----

## anotação
1. Como seu programa trata caracteres diferentes/acentuados.
 em primeiro momento usei o fmt.scan e não ajudou em lá essas coisas, pq espaços não era contados e eu queria que contasse. ai fui testando outros como o Scanln, depois o Scanf e ver (com testes) que satisfazia o que eu estava pensando na hora

estudando( coisa rápida)

vi que tem o

```go

bufio.NewScanner
```

acho que o scan serve para entrada de dados mas de uma maneira diferente da que pensei, parece que ele tem um foco em contagem, numeros

estudando... pode configurá-lo para ler palavras ou bytes


2. O que acontece em caso de empate.

- no meu programa em caso de empate está retornando o primeiro caractere, preciso trocar "letras", mas o ponto que notei 'com ajuda' é que o empate só acontece para contadores maiores que 1, no contrário vem um default, pq tirei o valor contado 1, e deixei os outros? acho que não pensei nisso.

em caso de empate: retornar os dois caracteres empatados


3. O que são aqueles ^[[A / ^[[D.

acho que pela forma que escrevi o código, ele lê cada movimento do curson como entrada, por isso aparece, por exemplo esses ai são setas para cima e para baixo, aaaaaaaaaaaaaaaaaaa, vo te que testa...
# testes - alter
…/dev/treinamento ✗ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
Digite sua mensagem: Digite sua mensagem: cas
Você digitou: cas
tamanho: 3
escreva uma palavra
^[[A
Todas as letras aparecem uma vez 1

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
Digite sua mensagem: Digite sua mensagem:    
Você digitou:    
tamanho: 3
escreva uma palavra
^[[A
Todas as letras aparecem uma vez 1

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
Digite sua mensagem: Digite sua mensagem: ^[[A
tamanho: 3ou: 
escreva uma palavra
^[[A
Todas as letras aparecem uma vez 1

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/main.go 
Digite sua mensagem: Digite sua mensagem:    ##
Você digitou:    ##
tamanho: 5
escreva uma palavra

"O que realmente acontece entre eu apertar uma tecla e fmt.Scan receber alguma coisa?"