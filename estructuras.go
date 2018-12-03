package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//Arrays:

	var arr1 [100]int

	arr2 := [3]string{"st1", "st2", "st3"}

	var arr3 [2]bool
	arr3[0] = true
	arr3[1] = false

	//Slices:

	//Con make se crea un array, pero modificable en TE.
	slc1 := make([]int, 10)

	//Añade espacios
	slc1 := append(slc1, 5)

	//Borrar el primer elemento
	slc1 := slc1[1:]

	//Borrar el segundo elemento
	slc1 = append(slc1[:1], slc1[2:])

	//Mapas:
	mapa := make(map[string]int)
	mapa["uno"] = 1
	mapa["dos"] = 2
	fmt.Println(mapa["uno"])

	//Parsear JSON a mapa:
	myJsonMap := make(map[string]interface{})
	jsonData := []byte(`{"hello":"world"}`)
	err := json.Unmarshal(jsonData, &myJsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", myJsonMap["hello"])

}
