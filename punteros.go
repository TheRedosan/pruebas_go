package main

func main() {

	numero := 5

	//Se guarda la referencia de memoria.
	pointer_numero := &numero

	//Para obtener la ref. de memoria:
	println(pointer_numero)

	//Para obneter el contenido:
	println(*pointer_numero)

}
