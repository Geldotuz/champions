package vigilante

import (
	"fmt"
	"time"
)

type Chequear interface {
	CheckCameras()
	AddVisitor(person string) bool
	Sleep()
}

// Cualquier struct que implemente los metodos de la Interfaz (TODOS E IGUALES) ya es esa interfaz

type Vigilant struct { // => Tambien es una Vigilancia
	Awake    bool
	Name     string
	Visits   int
	Building string
}

type Person struct {
	Name        string
	Age         int
	Nacionality string
	Sex         string
	Life        bool
}

// CountVisits return the numbers of visits
func CountVisits(v *Vigilant) int { //func that uses a vigilat
	v.Visits += 1
	return v.Visits
}

// called from main
func NewVigilant(awake bool, name string, visits int, building string) Chequear { // func that initializes a new vigilant
	return &Vigilant{
		Awake:    awake,
		Name:     name,
		Visits:   visits,
		Building: building,
	}
}
func NewPerson(name string, age int, nacionality string, sex string, life bool) *Person {
	return &Person{
		Name:        name,
		Age:         age,
		Nacionality: nacionality,
		Sex:         sex,
		Life:        life,
	}
}

func (p *Person) Dead() {
	p.Life = false
}

// Methods
func (v Vigilant) CheckCameras() {
	fmt.Println("I checked the cameras at", time.Now().String())
}
func (v Vigilant) AddVisitor(person string) bool {
	v.Name = "Hugo Rojas" // NEVER DO THIS
	fmt.Printf("I %s have registered %s\n", v.Name, person)
	return true
}

func (v *Vigilant) Sleep() {
	v.Awake = false
}

func (v *Vigilant) Respirar() {
	// NO IMPORTA LA LOGICA O IMPLEMENTATION
	fmt.Println("Estoy respirando")
}
