package players

func Aux(name string) *string {
	return &name
}

// 1  crea nuevo struct de un jugador de futbol. Usa nombres reales.
// Los campos que tiene el jugador son: Nombre, Apellido, Nacionalidad,  Si esta a prestamo (OnLoan),  Equipo Actual (string),
//  Equipo anterior (puntero string), goles, asistencias, posicion

type Player struct {
	Name         string
	Lastname     string
	Nacionality  string
	OnLoan       bool
	ActualTeam   string
	PreviousTeam *string
	Goals        int32
	Assists      int32
	Position     string
}

// 4 Crea una funcion que genere un nuevo jugador y retorne un puntero de ese jugador. Usa la funcion auxiliar para el campo que te de problemas

func NewPlayer(name string, lastName string, nacionality string, onloan bool, actualTeam string, previousTeam *string, goals int32, assists int32, position string) *Player {
	return &Player{
		Name:         name,
		Lastname:     lastName,
		Nacionality:  nacionality,
		OnLoan:       onloan,
		ActualTeam:   actualTeam,
		PreviousTeam: previousTeam,
		Goals:        goals,
		Assists:      assists,
		Position:     position,
	}
}
