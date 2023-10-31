package main

import (
	"champions-tournament/api"
	"champions-tournament/championsleague"
	"champions-tournament/players"
	"champions-tournament/premierleague"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("mensaje")
}

// NO BORRAR NINGUN COMENTARIO PARA ENTENDER EL ORDEN DE LAS COSAS
func main() {

	r := mux.NewRouter()

	a := &api.API{}
	a.RegisterRoutes(r)

	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	log.Println("listening")
	srv.ListenAndServe()

	// 6 crea un equipo nuevo de tu preferencia. Con datos reales

	zamora := premierleague.NewTeam("Zamora", "Agustin Tovar", 32000.45, "3")

	// 7 crea 1 nuevo jugador con data real para el equipo de tu preferencia. con 5 Goles y 3 Asistencias. No esta a prestamos
	miguel := players.NewPlayer("Miguel", "Celis", "Venezolano", false, "Zamora", players.Aux("Bayern"), 5, 3, "Forward")

	// 8 crea 1 nuevo jugador con data real para el equipo de tu preferencia.  Este jugador con un equipo anterior. Con 14 goles y 8 asistencias. No esta a prestamos
	oliver := players.NewPlayer("Oliver", "Beckand", "England", false, "Real Madrid", players.Aux("Barcelona"), 14, 8, "Center")

	// 9 crea un nuevo jugador con data real  para el equipo de tu preferencia. Este jugador a prestamo, 8 goles y 5 asistencia
	lionel := players.NewPlayer("Liones", "Messi", "Argentino", true, "Psg", players.Aux("Barcelona"), 14, 8, "Forward")

	// 9.1 crea un nuevo jugador con data real  para el equipo de tu preferencia. Con 10 Goles y 2 asistencias. No esta a prestamos
	manuel := players.NewPlayer("Manuel", "Muller", "Germany", false, "Bayern", players.Aux("Dormunt"), 10, 2, "Center")

	//NOTA: no uses variables intermedias aqui en el main, solo llama a la funcion y sus parametros. La posicion y Nacionalidad ponle el valor de tu preferencia

	// 10 Registra a los 4 jugadores en tu equipo

	zamora.RegisterPlayer(*miguel)
	zamora.RegisterPlayer(*oliver)
	zamora.RegisterPlayer(*lionel)
	zamora.RegisterPlayer(*manuel)

	// Tu equipo de la premier clasifico!! a Champions!!
	// 14 Crea tu equipo usando la funcion 11 para la champions ignorando a los jugadores
	CopaDelBurro := championsleague.NewChampionsLeagueTeam("1", zamora.Name)
	// 15 Llama al metodo #12 y registra a tus jugadores de la champions
	CopaDelBurro.RegisterPlayerInChampions(*zamora)
	// 16 Llamada al metdo para listar los jugadores registrados que jugaran la champions esta temporada
	CopaDelBurro.ListPlayersChamp()
}
