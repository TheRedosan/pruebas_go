package main

import (
	"errors"
	"fmt"
)

func main() {

	//Función anonima
	add := func(n int) int { return n + 1 }

	resultado := add(6)

	println(resultado)

	//Función clausura
	addN := func(m int) int {
		return func(n int) int {
			return m + n
		}
	}

	addFive := addN(5)
	result := addN(6)
	//5 + 6 must print 7

	println(result)

	err := doesReturnError()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", sum(1, 2, 3))
	fmt.Printf("%d\n", sum(4, 5, 6, 7, 8))

}

//Función
func bienv(nombre string) error {
	fmt.Printf("Bienvenido %s .\n")
	return nil
}

func doesReturnError() error {
	err := errors.New("this function simply returns an error")
	return err
}

func sum(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return
}
