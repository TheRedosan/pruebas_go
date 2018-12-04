package main

import "fmt"

//INTERFAZ
type AnchoVia interface {
	ComprobarAnchoVia() int
}

//Struct VIA ----------------------
type Via struct {
	Ancho int
}

func (v *Via) ViaTransitable(r Ancho) bool {
	return r.ComprobarAnchoVia() != v.Ancho
}

//Struct TREN ---------------------
type Tren struct {
	Ancho int
}

func (t *Tren) ComprobarAnchoVia() int {
	return t.Ancho
}

//MAIN -----------------------------
func main() {
	via := Via{Ancho: 10}

	passengerTrain := Tren{Ancho: 10}
	cargoTrain := Tren{Ancho: 15}

	canPassengerTrainPass := via.ViaTransitable(passengerTrain)
	canCargoTrainPass := via.ViaTransitable(cargoTrain)

	fmt.Printf("Puede pasar el tren de pasajeros? %b\n", canPassengerTrainPass)
	fmt.Printf("Puede pasar el tren de mercancías? %b\n", canCargoTrainPass)
}
