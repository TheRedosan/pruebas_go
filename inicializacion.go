package main

import (
	"errors"
	"log"
)

/*
func main() {

	var booleano bool     //False
	var entero int        //0
	var decimales float32 //0.0
	var cadena string     //""

	res := divisible(10, 0)

	fmt.Printf("%v\n", res)

}

func divisible(n, divisor int) bool {

	if divisor == 0 {
		//No se puede dividir entre 0.
		return false
	}

	return (n%divisor == 0)
}
*/

func main() {

	res, err := divisible(10, 0)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", res)
}

func divisible(n, divisor int) (bool, error) {
	if divisor == 0 {
		//No se puede dividir entre 0.
		return false, errors.New("A number cannot be divided by zero")
	}

	return (n%divisor == 0), nil
}
