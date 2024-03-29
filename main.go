package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Tch1b0/Esproso/brain"
	GameData "github.com/Tch1b0/Esproso/data"
)

var (
	games map[string]GameData.Data = make(map[string]GameData.Data)
)

type CustomSnake struct {
	Apiversion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}
type Turn struct {
	Move  string `json:"move"`
	Shout string `json:"shout"`
}

func main() {
	http.HandleFunc("/", customization)
	http.HandleFunc("/start", start)
	http.HandleFunc("/move", move)
	http.HandleFunc("/end", end)
	http.ListenAndServe(":5001", nil)
}

func customization(res http.ResponseWriter, req *http.Request) {
	snake := CustomSnake{
		Apiversion: "1",
		Author:     "Tch1b0",
		Color:      "#6100d5",
		Head:       "caffeine",
		Tail:       "coffee",
		Version:    "1",
	}

	json.NewEncoder(res).Encode(snake)
}

func start(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		return
	}

	CreateOrUpdateGame(req)

	res.WriteHeader(200)
	fmt.Fprint(res, "Ok")
}

func move(res http.ResponseWriter, req *http.Request) {
	start_time := time.Now()
	res.WriteHeader(200)
	data := CreateOrUpdateGame(req)

	next_move := brain.Move(data)

	json.NewEncoder(res).Encode(Turn{
		Move:  next_move,
		Shout: fmt.Sprintf("I am absolutely NOT going to move %s!", next_move),
	})
	fmt.Printf("MOVE request took %v\n", time.Since(start_time))
}

func end(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	CreateOrUpdateGame(req)
}

func CreateOrUpdateGame(req *http.Request) GameData.Data {
	data := GameData.Data{}
	json.NewDecoder(req.Body).Decode(&data)
	games[data.Game.Id] = data
	return data
}
