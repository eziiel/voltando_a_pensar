package main

import (
	"fmt"
	"strings"
)


func dark() {
   texto := "()})"
 
	 textoArray := strings.Split(texto, "")
	
	 
	 var validator []string

	 var valid = true

	 for _, c := range textoArray {

		if(c == "(") {
			validator = append(validator, c)
		} else if(c == "[") {
			validator = append(validator, c)
		} else if(c == "{") {
			validator = append(validator, c)
		} else if(c == ")") {
			if len(validator) == 0 {
				valid = false
				break
			}
			if(validator[len(validator)-1] == "(") {
				validator = validator[0:len(validator)-1]
			} else {
				valid = false
				break
			}
		} else if(c == "]") {
			if len(validator) == 0 {
				valid = false
				break
			}
			if(validator[len(validator)-1] == "[") {
				validator = validator[0:len(validator)-1]
			} else {
				valid = false
				break
			}
		} else if(c == "}") {
			if len(validator) == 0 {
				valid = false
				break
			}
			if(validator[len(validator)-1] == "{") {
				validator = validator[0:len(validator)-1]
			} else {
				valid = false 
				break 
			}
		}
	 }
	 if len(validator) != 0 {
		valid = false
	 }
	 fmt.Println(valid)
	//  for i, c := range texto{
	// 	fmt.Println("------------------------------------------------------------", validator, "c: ", c)
	// 	 if (len(validator) == 0 && i == 0) {
	// 		validator = append(validator, int(c))
	// 		} 		

	// 		fmt.Println("len(validator): ",len(validator), "i: ", i, "len(texto): ", len(texto), "len(texto) - 2", len(texto))
	// 		if len(validator) == 0 && i == len(texto) {
	// 			validator = append(validator, int(c))
	// 			break
	// 	 }
			
	// 	//  fmt.Println("i > 0: ", i > 0 )
	// 	fmt.Println("i : ", i, "c: ", c, validator)
	// 	 if i > 0 {
	// 		 	//fmt.Println("i > 0 : ", i > 0 )
	// 		// 	fmt.Println("c: ", c , rune(validator[i-1]), c - rune(validator[i-1]))
  //     if(len(validator) == 0 && i != 0) {
  //       validator = append(validator, int(c))
  //     }

	// 			var count = math.Abs(float64(int(c) - validator[len(validator) - 1]))
	// 			fmt.Println("validator5: ", validator, i, c, validator[len(validator) - 1])

	// 			fmt.Println("count",  count, "<= 5", count <= 5, "> 0", count > 0 ,  "<= 5 &&count >= 0 ", count <= 5 &&count >= 0 )

	// 			if (count <= 5 && count > 0 ) {
	// 				fmt.Println("validator1111111111111111: ",)
	// 				validator = validator[0:len(validator)-1]
	// 				fmt.Println("validator33333333333333333333333: ", validator)
	// 			} else {
	// 				validator = append(validator, int(c))
	// 			}

	// 		 } 
	// 	}
	// 	fmt.Println("validator5: ", validator)
	}