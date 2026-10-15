## Problema

🧠 Exercício 02 — Primeiro caractere não repetido

Crie uma função que receba uma string e retorne o primeiro caractere que aparece apenas uma vez.

Exemplos:

```
"banana"
→ "b"
```

```
"abacaxi"
→ "b"
```

```
"aabbc"
→ "c"
```


```
"aabb"
→ nenhum
```


Mas tem uma pegadinha importante:

```
"swiss"
```

As frequências são:

```
s → 3
w → 1
i → 1
```

O resultado é:

```
w
```


Porque queremos o primeiro na ordem em que aparece, não simplesmente qualquer caractere com frequência 1.
...

## Minha interpretação

vou ter que usar mais de um loop novamente, pensando se tem alguma função em go para fazer isso, no JS/TS algo como lastIndexOf firstIndexOf, mas aqui não sei se existe, de todo modo quero treinar minha análise e usar algo que não seja óbvio

talvez dois loops, acho que variáveis para guardar valores



...

## Estratégia

- uma entrada de dados
- transformar em array
- um loop para percorrer cada item da string, salvando o item atual para comparaçãoy
- dentro dele um outro loop para comparar o item atual com o item do loop concatenado acima
- duas váriáveis fora dos loops, uma para salvar o item que aparece menos e outra para contar o total de vezes que ele aparece
- mas isso seria igual a solução antiga, mas ao contrário


 - como vou ter vou ter um contador e salvar a letra em validação vou já supor que elá tem um item = int = 1, sendo assim
 vou validar se o count dele é 2, se for é o primeiro, já retorna ele e return para finalizar o loop

--------------------

- e se e passasse pelo array e criasse uma struct {'a': 1, 'b':2, 'c': 2, ...} e retornasse todos que são um (não pediu isso...)
- depois 
...

- levando em consideração que vai ser o primeiro item, os demais eu nem precisao validar
- mas e se o ultimo for o que aparece menos então não importa
- vou fazer comparação de valores

...

## Estrutura de dados

...

## Complexidade esperada

- O(n²) na minha cabeça, não sei se está certo, digo isso pq é um loop dentro de outro, foda que não sei uma alternativa para isso :/

On é o primeiro, e para cada On vou ter um On, então On * On = O², nessa ideia seria O(n^2) ou O²n? não sei explicar muito bem ._. 

...

## Casos de teste

- string vazia, emoji, acentos, palindromos


...

## comentários

- a ideia de: não pediu então não faço me pegou um pouco pq eu tento ser pro ativo, por exemplo:
- em caso de empate
- em caso que não tenha items que não se repete
- pensei em caracteres matemáticos:


```
∃: Lê-se "existe". Por exemplo, ∃ x significa "existe um x".
∃!: Lê-se "existe um único" (um valor exato que cumpre a condição).
∄: Lê-se "não existe"
```

nada pedido....

'

## comentários 2

tive um problema, vou tentar entender:
logs:
```
…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: csssss
csssss
c
a letra:  c  é a primeira que aparece somente uma vez!

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: swisswi  
swisswi
s
w
a letra:  w  é a primeira que aparece somente uma vez!

```

descobri, eu estava resetando o contador para 0, ksaksak, mas eu disse que era para 1, descobri rápido e sem uso de prints, agora uma dúvida, e se tiver items que não se repetem?

resolvido também, rápidamente, mas penso: óbivio isso, e demorei (menos de um minuto), mas o rolê de fazer o que não pediu entra em: como não pensei nisso, fico pensando em coisas óbivias que não pensei...


```
…/dev/treinamento ✗ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: casa
casa
a letra:  c  é a primeira que aparece somente uma vez!

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: swiss
swiss
a letra:  w  é a primeira que aparece somente uma vez!

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: swissw
swissw
a letra:  i  é a primeira que aparece somente uma vez!

…/dev/treinamento ❯ swisswi
bash: command not found: swisswi

Não existe itens que se repetem apenas uma vez!na_01/exercicios/dia_02/main.go 
Digite sua mensagem: swisswi^[[A^[[A^[[A
…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: casa
casa
a letra:  c  é a primeira que aparece somente uma vez!

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: casac
casac
a letra:  s  é a primeira que aparece somente uma vez!

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: casacs
casacs
Não existe itens que se repetem apenas uma vez!

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
Digite sua mensagem: ^[[C^[[A
a letra:  C  é a primeira que aparece somente uma vez!

a letra:  D  é a primeira que aparece somente uma vez!xercicios/dia_02/main.go 
Digite sua mensagem: ^[[A^[[A^[[A^[[C^[[C^[[D

```