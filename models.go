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

type PlayerStatsYear {
  Id int64 `json:"id"`
  Stint string `json:"stint"`
  TeamId string `json:team_id`
  LeagueId string `json:league_id`
  GamesPlayed int64 `json:games_played`
  GamesStarted int64 `json:games_started`
  Minutes int64 `json:minutes`
  Points int64 `json:points`
  ORebounds int64 `json:o_rebounds`
  DRebounds int64 `json:d_rebounds`
  Rebounds int64 `json:rebounds`
  Assists int64 `json:assists`
  Steals int64 `json:steals`
  Blocks int64 `json:blocks`
}
