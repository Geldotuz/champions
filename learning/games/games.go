package games

import (
	"errors"
	"fmt"
)

type Game struct {
	Name  string
	Price float64
}

type Juego struct {
	Price       float64
	Plataform   string
	Descripcion string
}

type Report struct {
	Fecha     string
	Caducidad string
	Juego     Game
	AreNul    bool // true si ambos son nulos
}

func ReciveMap(juegos map[string]Juego) error {
	for name, juego := range juegos {
		fmt.Println(name, juego)
	}
	return nil
}

func Print(games []Game) {
	for i := 0; i <= len(games)-1; i++ {
		g := games[i]
		fmt.Println(g)
	}
	for _, game := range games {
		fmt.Println(game)
	}
}

func Receive(g1, g2 *Game) (Report, error) {

	r := Report{
		Fecha:     "23 de Mayo",
		Caducidad: "23 de Junio",
		AreNul:    false,
		Juego:     Game{},
	}

	if g1 == nil && g2 == nil {
		r.AreNul = true
		return r, errors.New("both games are nil")
	}

	if g1 == nil || g2 == nil {
		return r, errors.New("one game ames are nil")
	}

	if g2.Price == g1.Price {
		return r, errors.New("the games cost the same amount")
	}

	if g1.Price > g2.Price {
		r.Juego = *g1
		return r, nil
	}

	r.Juego = *g2
	// g2->Direccion *g2-> Valor de lo que tiene el puntero
	return r, nil

}
