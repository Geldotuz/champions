package premierleague

import "champions-tournament/players"

// 2 crea nuevo struct de un equipo de la premier league. Usa nombres reales.
//Los campos que tiene el equipo son: Nombre, Estadio, Capital(Dinero), Jugadores (array), Temporada Actual (string)

type Team struct {
	Name          string
	Stadium       string
	Capital       float64
	Players       []players.Player
	CurrentSeason string
}

// 3 Crea una nueva funcion que genera un nuevo equipo de futbol y retorne un puntero del equipo nuevo. El campo jugadores puede estar vacio

func NewTeam(name string, stadium string, capital float64, currentSeason string) *Team {
	return &Team{
		Name:          name,
		Stadium:       stadium,
		Capital:       capital,
		CurrentSeason: currentSeason,
	}
}

// 5 Crea un nuevo metodo para ingresar/registrar jugadores a la plantilla del equipo

func (t *Team) RegisterPlayer(p players.Player) {
	t.Players = append(t.Players, p)
}
