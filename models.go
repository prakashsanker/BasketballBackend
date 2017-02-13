package main


type Player struct {
  Id int64 `json:"id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Position string `json:"position"`
  Height int64 `json:"height"`
  Weight int64 `json:"weight"`
  PlayerId string `json:"player_id"`
}

type Players []Player
