package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Delcaración explicita
	var nombre1 string = "Álvaro"

	//Delcaración implicita
	var nombre2 = " Redondo "

	//Delcaración inferred
	nombre3 := "Sánchez"

	//Impresion
	println(nombre1 + nombre2 + nombre3)

	//Impresion de los tipos de variable
	fmt.Println("La variable 'nombre1' tiene formato:", reflect.TypeOf(nombre1))
	fmt.Println("La variable 'nombre3' tiene formato:", reflect.TypeOf(nombre3))
}
