package main

import (
	"fmt"
)

func main() {
  texto := "((()))}"

  var pares = map[rune]rune{
    ')': '(',
    '}': '{',
    ']': '[',
  }

  var count []string

  var validator = true

  i := 0
  
  for i < len(texto) {
    
    par, existe := pares[rune(texto[i])]
    
    fmt.Println("par, existe", string(par), existe)


    if !existe {
      count = append(count, string(texto[i]))
    } else {
      if len(count) == 0 {
        validator = false
        break
      }

       if string(par) == count[len(count)-1] {
        count = count[0:len(count)-1]
      } else {
        validator = false
        break
      }
    }
      
    i++
  }
  if len(count) != 0 {
    validator = false
}

  fmt.Println(validator, count)
}