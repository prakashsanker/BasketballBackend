package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WELCOME!")
}


func searchPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	query := "SELECT id, player_id, first_name, last_name, position, height, weight from player_identity WHERE first_name LIKE" + "'%" + query + "%'"
	rows, err := db.Query(query)
	check(err)
	var player Player
	var players Players
	for rows.Next() {
		var playerId string
		var firstName string
		var lastName string
		var position string
		var height int64
		var weight int64
		var id int64
		err := rows.Scan(&id, &playerId, &firstName, &lastName, &position, &height, &weight)
		check(err)
		player = Player{Id: id, FirstName: firstName, LastName: lastName, PlayerId: playerId, Position: position, Height: height, Weight: weight}
		players = append(players, player)
	}

	if err := json.NewEncoder(w).Encode(players); err != nil {
		check(err)
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
