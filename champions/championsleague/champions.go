package championsleague

import (
	"champions-tournament/premierleague"
	"fmt"
)

const champ = "CHAMP-"

type ChampionsPlayer struct {
	Name     string
	Lastname string
	Team     string
	Position string
}

type ChampionsLeagueTeam struct {
	Season            string
	TeamName          string
	RegisteredPlayers []ChampionsPlayer
}

// 11 crea una funcion que genere un nuevo equipo de la champions league. El campo de los jugadores registrados vacio.
func NewChampionsLeagueTeam(season string, teamName string) *ChampionsLeagueTeam {
	return &ChampionsLeagueTeam{
		Season:   season,
		TeamName: teamName,
	}
}

// 12 Crea un metodo RegisterPlayerInChampions que reciba un equipo de futbol de la premier. Y registre en la champions a sus jugadores con las siguientes reglas:
func (c *ChampionsLeagueTeam) RegisterPlayerInChampions(team premierleague.Team) {
	// * Jugador a prestamo no se registra
	for _, player := range team.Players {
		if player.OnLoan {
			continue
		}
		// * Jugador con bajo rendimiento, queda fuera. Bajo Rendimiento es menos de 6 goles o menos de 4 asistencias
		if player.Goals < 6 && player.Assists < 4 {
			continue
		}
		pl := ChampionsPlayer{
			Name:     fmt.Sprintf("%s%s", champ, player.Name),
			Lastname: fmt.Sprintf("%s%s", champ, player.Lastname),
			Team:     champ + player.ActualTeam,
			Position: champ + player.Position,
		}
		c.RegisteredPlayers = append(c.RegisteredPlayers, pl)
	}
}

// * Todos los campos del jugador registrado en la champions debe llevar el prefijo CHAM-
// 13 Crea un metodo que liste los jugadores inscritos en la champions y se imprima todos los campos de los jugadores usando el fmt.Printf

func (c *ChampionsLeagueTeam) ListPlayersChamp() {
	for _, player := range c.RegisteredPlayers {
		fmt.Printf("Jugador: %+v \n", player)
	}
}
