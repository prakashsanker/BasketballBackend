package main
import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "io/ioutil"
  "strings"
)

func main() {
  db,err := sql.Open("mysql", "root:prefontainetqcio123@/basketball")
  err = db.Ping()

  if err != nil {
    fmt.Println("Failed to prepare connection to database")
    log.Fatal("Error:", err.Error())
  }

  defer db.Close()

  content, err := ioutil.ReadFile("./dataEntry/basketball_master.csv")
  check(err)

  lines := strings.Split(string(content), "\r")

  for _, line := range lines {
    line = strings.TrimSpace(line)
    splitStr := strings.Split(line, ",")
    _, err := db.Exec("INSERT INTO player_identity(player_id, first_name, last_name, position, height, weight, college, birth_date) VALUES(?,?,?,?,?,?,?,?)", splitStr[0], splitStr[1], splitStr[2], splitStr[3], splitStr[4], splitStr[5], splitStr[6], splitStr[7])
    check(err)
  }

  content, err = ioutil.ReadFile("./dataEntry/basketball_players.csv")
  check(err)

  lines = strings.Split(string(content), "\r")

  for _, line := range lines {
    line = strings.TrimSpace(line)
    splitStr := strings.Split(line, ",")
    values := "INSERT INTO players(player_id, year, stint, team_id, league_id, games_played, games_started, minutes, points, o_rebounds, d_rebounds, rebounds, assists, steals, blocks, turnovers, personal_fouls, fg_attempted, fg_made, ft_attempted, ft_made, three_attempted, three_made, post_games_played, post_games_started, post_minutes, post_points, post_o_rebounds, post_d_rebounds, post_rebounds, post_assists, post_steals, post_blocks, post_turnovers, post_pf, post_fg_attempts, post_fg_made, post_ft_attempts, post_ft_made, post_three_attempt, post_three_made) VALUES("
    for _, val := range splitStr {
      if val != "" {
        toAppend := "\""+ val + "\"" + ","
        values += toAppend
      }
    }
    values = strings.TrimSuffix(values, ",")
    values = strings.TrimSuffix(values, ",")
    values = values + ")"
    _, err := db.Exec(values)
    check(err)
  }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
