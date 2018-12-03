package main

import (
	"fmt"
)

func main() {

	numero := 10

	//IF
	if numero == 20 {
		println("La variable es 20. No aparece este mensaje")
	} else if numero != 10 {
		println("Si la variable no es 10.")
	} else {
		println("La variable es distinta de 20. Se imprime este")
	}

	//Switch
	switch numero {
	case 1:
		println("El numero es 1.")
	case 10:
		println("El numero es 10.")
	}

	//For
	for i := 0; i <= 10; i++ {
		println(i)
	}

	//Foreach
	array := [5]int{0, 1, 2, 3, 4}

	for index, valor := range array {
		fmt.Printf("Index: %d | Valor: %d \n", index, valor)
	}

}
