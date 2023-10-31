package main

import (
	"fmt"
	"martin/games"
	"martin/vigilante"
)

func main() {

	// var array []games.Game

	game1 := games.Juego{
		Price:       59.99,
		Plataform:   "Switch",
		Descripcion: "the game is very good",
	}
	game2 := games.Juego{
		Price:       4.38,
		Plataform:   "Gba",
		Descripcion: "A retro game",
	}
	game3 := games.Juego{
		Price:       99.99,
		Plataform:   "Ps5",
		Descripcion: "This game is new the price is in 99.99",
	}

	juegos := make(map[string]games.Juego)

	juegos["MarioBross"] = game1
	juegos["Diddykong"] = game2
	juegos["BobEsponja"] = game3

	err := games.ReciveMap(juegos)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	juanito := vigilante.NewVigilant(true, "Juanito", 20, "Megaron")
	juanito.Sleep()
	juanito.AddVisitor("")
	juanito.CheckCameras()
	/*
		visits := vigilante.CountVisits(juanito)
		human := vigilante.NewPerson("angel", 30, "Venezuela", "Male", true)
		human.Dead()
		fmt.Println(human)
		fmt.Println()
		fmt.Println(juanito.Name, " registro ", visits, " visitas")
		fmt.Println(juanito.Name, " registro ", juanito.Visits, " visitas")
		juanito.CheckCameras()
		vino := juanito.AddVisitor("Angel Martin")
		juanito.Sleep()
		fmt.Println(juanito)
		_ = vino*/
}
