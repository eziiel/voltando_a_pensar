# Exercício 03 — Parênteses

## Problema
🧠 Exercício 03 — Validar parênteses
Receba uma string contendo parênteses e determine se eles estão corretamente balanceados.

Exemplos:

```
"()"       → válido
"(())"     → válido
"()()"     → válido
"(()())"   → válido
Inválidos:
```


```
"("        → inválido
")"        → inválido
"(()"      → inválido
"())"      → inválido
")("       → inválido
```

...

## Minha interpretação

bom, pensando de maneira rápida, se o item da string for "(" somo 1, se for ")" subtraio 1, no final de cada iteração eu validaria se o valor é 0 e aplico em um contador, para ver o total de pares... pensando aqui: poderia ter um somador para "(" e um para ")", no final, pela quantidade do somador eu descobriria o par.....

me obrigou a pensar: 


You said:
￼
main.go
File
## Problema

🧠 Exercício 02 — Primeiro caractere não repetido

Crie uma função que receba uma string e retorne o primeiro caractere que aparece apenas uma vez.

Exemplos:

"banana"
→ "b"


"abacaxi"
→ "b"


"aabbc"
→ "c"



"aabb"
→ nenhum



Mas tem uma pegadinha importante:

"swiss"


As frequências são:

s → 3
w → 1
i → 1


O resultado é:

w



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


∃: Lê-se "existe". Por exemplo, ∃ x significa "existe um x".
∃!: Lê-se "existe um único" (um valor exato que cumpre a condição).
∄: Lê-se "não existe"


nada pedido....

'

## comentários 2

tive um problema, vou tentar entender:
logs:
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


descobri, eu estava resetando o contador para 0, ksaksak, mas eu disse que era para 1, descobri rápido e sem uso de prints, agora uma dúvida, e se tiver items que não se repetem?

resolvido também, rápidamente, mas penso: óbivio isso, e demorei (menos de um minuto), mas o rolê de fazer o que não pediu entra em: como não pensei nisso, fico pensando em coisas óbivias que não pensei...


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
￼
Show more
￼
￼
￼

ChatGPT said:
KKKK você cumpriu a disciplina. Agora deixa eu corrigir algumas coisas do raciocínio, porque tem umas pequenas imprecisões que valem ouro nesse treino.

1. A conclusão principal está certa
Você chegou em:

"eu não preciso aninhar, ainda uso dois loops"

Exatamente.

E isso é uma distinção importante:

￼
2 loops ≠ O(n²)
O que importa é como eles se relacionam.

￼
Loop 1: n
Loop 2: n

n + n = 2n → O(n)
Enquanto:

￼
Loop 1: n
    Loop 2: n

n × n = n² → O(n²)
O(1)
O(log n)
O(n)
O(n log n)
O(n²)
O(2ⁿ)
10
20
30
40
50
60
50
100
150
200
Tamanho da entrada n
Oper.
𝑛
=
16
:
(
1
,
log
⁡
2
𝑛
,
𝑛
)
=
(
1
,
4
,
16
)
n=16:(1,log 
2
​
 n,n)=(1,4,16)
(
𝑛
log
⁡
2
𝑛
,
𝑛
2
)
=
(
64
,
256
)
(nlog 
2
​
 n,n 
2
 )=(64,256)
2
𝑛
=
65.536
2 
n
 =65.536
𝑛
n
￼
4
￼
8
￼
16
￼
32
￼
64
4
8
16
32
64
Escolha um tamanho de entrada para comparar o crescimento de seis classes comuns de complexidade de tempo. A guia vertical e as contagens de operações são atualizadas juntas; uma seta para cima indica que a curva continua além do gráfico.
￼
Dar feedback
2. Tem um errinho na sua contagem
Você escreveu:

"passo por cada item n-1"

e:

"para cada item percorro o range de n-2"

Não precisa pensar assim.

Se existem 5 caracteres, o algoritmo pode fazer aproximadamente 5 operações na primeira passagem e aproximadamente 5 na segunda:

￼
5 + 5 = 10
Não:

￼
5² = 25
E mesmo que fossem 4 + 4, 5 + 5, etc., a ordem de crescimento continua sendo linear.

É por isso que:

￼
O(2n) = O(n)
Os termos constantes são ignorados na notação Big O.

3. E cuidado com "percorro a struct"
Aqui tem uma sutileza.

Você não está necessariamente percorrendo o map na segunda etapa.

Na sua implementação, você faz:

Go
￼
for i := range palavra {
    if letterCount[letters[i]] == 1 {
        ...
    }
}
Ou seja, você percorre a palavra novamente e consulta o map.

Isso é melhor para a especificação porque preserva a ordem original:

￼
palavra:
s w i s s

map:
s → 3
w → 1
i → 1
Se você simplesmente percorresse o map, não poderia usar a ordem dele para determinar "o primeiro", porque map não é uma estrutura para preservar essa ordem.

E isso é uma descoberta interessante para registrar.

4. Sua melhor frase foi esta
"de uma forma, ao meu ver, mais limpa, menos propensa a erros e mais fácil de debugar."

Sim.

Mas acrescentaria:

e com uma complexidade assintótica melhor.

Você não apenas deixou o código "mais bonito".

Você mudou o algoritmo.

De:

￼
para cada elemento:
    procure em todos os elementos
para:

￼
conte todos
depois procure o primeiro único
Isso é uma mudança de estratégia.

5. E agora quero que você perceba uma coisa sobre você mesmo
Você começou esse exercício dizendo:

"to com preguiça de fazer isso"

Depois escreveu a estratégia.

Depois teve bug.

Depois corrigiu.

Depois encontrou uma segunda solução.

Depois percebeu que não precisava aninhar loops.

Isso é justamente o oposto daquela sensação de:

"eu só sei programar quando a IA me dá o caminho."

Você acabou de produzir o caminho.

Não romantiza isso — você ainda está enferrujado mesmo. 😂
Mas enferrujado é diferente de não saber.

E esse exercício já cumpriu seu objetivo.

Antes de encerrarmos o Dia 2
Só falta aquela questão que ficou pendurada:

range palavra vs range letters + Unicode.

Não precisa resolver agora se você estiver cansado. Mas esse é o próximo ponto que eu quero que você investigue, porque conecta diretamente com o que você anotou ontem sobre string, byte e rune.

E aí fechamos a semana com uma coisa muito boa: você já saiu de "como faço esse loop?" para começar a perguntar "qual representação e qual estratégia fazem sentido para esse problema?".

￼
￼
￼
￼
￼

You said:
…/dev/treinamento ❯ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
çé😀
primeira contagem:  map[]
panic: runtime error: index out of range [4] with length 3

goroutine 1 [running]:
main.main()
        /home/eziel/Documents/dev/treinamento/ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go:60 +0x3e5
exit status 2

…/dev/treinamento ✗ go run ciclos/ciclo_01/semana_01/exercicios/dia_02/main.go 
çé😀
primeira contagem:  map[]
segunda contagem:  map[ç:1 é:1 😀:1]
a letra:  ç  é a primeira que aparece somente uma vez!

…/dev/treinamento ❯ 

primero com range em letterCount(map) e outro em letters string.split  pq?
￼
Show more
￼
￼
￼

ChatGPT said:
Exatamente — você acabou de encontrar na prática a diferença entre bytes e caracteres Unicode em Go. E o seu erro é uma ótima demonstração disso.

Você testou:

￼
çé😀
e teve:

￼
panic: index out of range [4] with length 3
O que aconteceu?
Você tem:

Go
￼
letters := strings.Split(palavra, "")
Para:

￼
çé😀
o letters fica, conceitualmente:

￼
índice    valor
  0        ç
  1        é
  2        😀
Portanto:

￼
len(letters) = 3
Mas quando você faz:

Go
￼
for i := range palavra
você não está percorrendo os índices de letters.

Você está percorrendo os índices da string em bytes.

E esses caracteres ocupam quantidades diferentes de bytes em UTF-8:

￼
ç  → 2 bytes
é  → 2 bytes
😀 → 4 bytes
Então a string possui 8 bytes, embora tenha 3 caracteres Unicode (runes).

Por isso você acaba tentando algo equivalente a:

￼
letters[0] → ç  ✅
letters[1] → é  ✅
letters[2] → 😀  ✅
letters[3] → ??? ❌
letters só tem índices 0, 1 e 2.

Daí:

￼
index out of range [4] with length 3
E aí vem a parte interessante
Quando você mudou para percorrer letters, funcionou:

Go
￼
for i := range letters {
    letterCount[letters[i]]++
}
Porque agora os dois estão no mesmo "universo":

￼
letters

0 → ç
1 → é
2 → 😀
Mas por que range em string faz isso?
Porque Go trata string fundamentalmente como uma sequência de bytes.

Quando você faz:

Go
￼
for i := range palavra
o i representa a posição em bytes onde começa cada rune UTF-8.

Não é simplesmente:

￼
0, 1, 2, 3...
No seu exemplo, seria aproximadamente:

￼
byte:
0 1 | 2 3 | 4 5 6 7
 ↓     ↓       ↓
 ç     é       😀
O range sobre uma string, porém, tem uma característica importante: ele decodifica UTF-8 para produzir os runes.

Então existe uma diferença entre:

Go
￼
for i := range palavra
e:

Go
￼
for i, r := range palavra
No segundo caso:

Go
￼
r
é o rune/caractere Unicode.

E olha que interessante o que você fez
Ontem você escreveu:

"Diferença entre byte, rune e string."

Hoje você acabou de encontrar um bug causado exatamente por isso.

Isso é muito melhor do que eu simplesmente te explicar:

"Go usa UTF-8 e rune representa Unicode."

Você acabou de ver a consequência.

E guarda essa frase:

len(string) não significa necessariamente "quantidade de caracteres".

Para:

￼
casa
parece igual.

Para:

￼
çé😀
a diferença aparece.

Só uma correção no seu comentário
Você falou:

"primeiro com range em letterCount(map) e outro em letters string.split"

Na verdade, não foi range no letterCount no código que você mostrou.

Foi:

Go
￼
for i := range letters
nas duas passagens.

O map está sendo consultado:

Go
￼
letterCount[letters[i]]
Isso importa porque você está preservando a ordem da palavra/letters, enquanto o map serve como tabela de contagem.

E isso fecha muito bem o exercício:

￼
string
   ↓
strings.Split
   ↓
[]string
   ↓
primeira passagem → map[string]int
   ↓
segunda passagem → encontra primeiro count == 1
Agora sim eu considero o Dia 2 concluído.

E, sinceramente: esse panic com çé😀 foi provavelmente a melhor coisa que poderia ter acontecido nesse exercício. Você não só aprendeu que a solução O(n) funciona; encontrou um problema real de representação de dados no processo.

￼
￼
￼
￼
￼

Today 8:25 PM
You said:
vamos começar
￼
￼
￼

ChatGPT said:
BORA. 🔥

Vamos para o Dia 3 da Semana 01.

Até aqui você já trabalhou:

Dia 1 → contagem de frequência com O(n²)

Dia 2 → mesma ideia transformada em O(n) usando map

e ainda esbarrou em string/UTF-8/rune na prática

Hoje quero mudar o tipo de raciocínio. Menos contagem, mais decomposição.

🧠 Exercício 03 — Validar parênteses
Receba uma string contendo parênteses e determine se eles estão corretamente balanceados.

Exemplos:

￼
"()"       → válido
"(())"     → válido
"()()"     → válido
"(()())"   → válido
Inválidos:

￼
"("        → inválido
")"        → inválido
"(()"      → inválido
"())"      → inválido
")("       → inválido
Mas vamos colocar uma segunda camada:

￼
"(abc)"    → válido
"(a(b)c)"  → válido
"a(b)c"    → válido
Ou seja, outros caracteres podem existir na string. O que importa são os parênteses.

Antes de escrever código
Quero que você pense numa coisa:

Se eu estiver lendo:

```
"(()())"
```

em determinado momento, preciso saber alguma coisa sobre os parênteses que já apareceram.

O que exatamente?

Não quero que você responda "um contador" automaticamente.

Pensa no comportamento:

```
(
(
)
(
)
)
```
Quando aparece (, acontece alguma coisa.

Quando aparece ), acontece alguma coisa.

E existe uma condição que não pode acontecer durante o percurso.


---
pensei que parenteses são escopos então (()()) são três escopos, facil encontrar onde começam
e onde terminam, e se eu mostrar a posição do escopo que não esta completo? adivinhar onde ele fecha
não tem como, mas posso encontrar onde ele começa assim fica mais fácil para fechar ou ver se é 
válido ter... levando em consideração que pode haver outros caracteres no meio. "E se eu tiver viajando ou pensando a coisa errada..." 


...

## Estratégia
 - aplicar o contador primeiro
 - acho que vou ter que trabalhar com indeces 
 - depois ir para a posição do item(s:parenteses) que não se completam

...

## Estrutura de dados

- map
- for
- string
- int
- slice?
...

## Complexidade esperada
 - O(n) ? 
...

## Casos de teste
  - n caracteres diferentes 

...

## Dúvidas

 
...

## comentários durante o exercício
a ia me deu uma solução (auto complet), ai aceitei e fui entender, eu comentei, tinha condicional para validar se o loop terminou e comparadores de = == != !==, percebi, primeiro tirei os comparadores de dentro do loop, depois vi que tinha complexidade de maos (if desnecessário) e no fim resolvia apenas como contador e não a que eu pensei de segundo caso.

salvar em um array os indices dos "(" aberturas, corto o array

-----------------
## Estratégia

Percorrer a expressão mantendo uma pilha de parênteses abertos.

- `(` → entra na pilha
- `)` → remove o último `(` da pilha
- outros caracteres → ignora

## Estado que preciso manter

- Os parênteses `(` que ainda não foram fechados.

## Condição inválida durante o percurso

- Encontrar `)` quando a pilha está vazia.

## Condição no final

- Pilha vazia → expressão válida.
- Pilha com elementos → existem `(` sem `)`.

## histórico:
	for i := range palavras {
		var lettersCount = make(map[string]int)
		var letters = strings.Split(palavras[i], "")
		
		for i := range letters {
			lettersCount[letters[i]] = lettersCount[letters[i]] + 1
		}
	
		anagramas = append(anagramas, lettersCount)

	}
		
		fmt.Println("anagramas: ", anagramas)

---
