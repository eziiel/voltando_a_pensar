# Pensa:


Se eu pudesse saber a quantidade de cada caractere antes de procurar o primeiro único, eu precisaria mesmo de um loop dentro de outro?


## Comparação

to com preguiça de fazer isso, mas vou presar pela disciplina

### Estratégia 1
Como funciona:

a etratégia um é aninhamento de loops, acaba deixando mais complexo, não que esteja errado, mas é mais dificil de acompanhar e de debugar, e o tempo de execução é mais longo


eu passo por cada item n-1, e, para cada item eu percorro o range de n-2, então 5 letras O(5²) = 25 iteraçoes
...
Complexidade: O(n²)
...

### Estratégia 2
Como funciona:

uso dois loops mas fora um do outro, não aninhados, eu passo pelo length da palavra n-1, crio uma struct e depois eu percorro a struct, então 5 items do map O(5) = 5 iteraçoes no máx, ou até achar o item que eu preciso, em ordem crescente, decrescente seria o último
...
Complexidade: O(n)
...

### O que mudou no meu raciocínio?
...

eu não preciso aninhar, ainda uso dois loops, ainda seto variáveis dora dos loops, mas, de uma forma, ao meu ver, mais limpa, menos propensa a erros e mais fácil de debugar e com uma complexidade assintótica melhor.. 


Mas por que range em string faz isso?
Porque Go trata string fundamentalmente como uma sequência de bytes.

…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
çé😀
primeira contagem:  map[]
panic: runtime error: index out of range [4] with length 3

goroutine 1 [running]:
main.main()
        /home/eziel/Documents/dev/treinamento/ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go:60 +0x3e5
exit status 2
-----
E guarda essa frase:

len(string) não significa necessariamente "quantidade de caracteres".