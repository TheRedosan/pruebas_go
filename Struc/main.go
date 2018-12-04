package main

import "fmt"

func main() {

	p := Persona{
		Nombre:   "Álvaro",
		Apellido: "Redondo",
		Hobbies:  []string{"jugar videoguejos", "correr", "comer"},
		id:       "sa3-223-asd",
	}

	fmt.Printf("A %s le gusta %s, %s y %s\n", p.GetNombreCompleto(), p.Hobbies[0], p.Hobbies[1], p.Hobbies[2])

}

type Persona struct {
	Nombre   string
	Apellido string
	Hobbies  []string
	id       string
}

func (persona *Persona) GetNombreCompleto() string {
	return fmt.Sprintf("%s %s", persona.Nombre, persona.Apellido)
}
