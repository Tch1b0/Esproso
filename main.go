package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tch1b0/Esproso/brain"
	GameData "github.com/Tch1b0/Esproso/data"
)

var (
	games map[int]GameData.Data = make(map[int]GameData.Data )
)

type CustomSnake struct {
	Apiversion string `json:"apiversion"`
	Author 	   string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}
type turn struct {
	Move string `json:"move"`
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
	snake := CustomSnake {
		Apiversion: "1",
		Author: "Tch1b0",
		Color: "#888888",
		Head: "default",
		Tail: "default",
		Version: "0.1",
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
	res.WriteHeader(200)
	data := CreateOrUpdateGame(req)

	next_move := brain.Move(data)

	json.NewEncoder(res).Encode(turn {
		Move: next_move,
		Shout: fmt.Sprintf("I am absolutely NOT going to move %s!", next_move),
	})
}

func end(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	CreateOrUpdateGame(req)
}


func CreateOrUpdateGame(req *http.Request) GameData.Data {
	var game GameData.Game
	json.Unmarshal([]byte(req.FormValue("game")), &game)
	turn := req.FormValue("turn")
	var board GameData.Board
	json.Unmarshal([]byte(req.FormValue("board")), &board)
	var you GameData.Snake
	json.Unmarshal([]byte(req.FormValue("you")), &you)

	data := GameData.Data {
		Game: game,
		Turn: turn,
		Board: board,
		You: you, 
	}
	games[data.Game.Id] = data

	fmt.Println(data.Board.Food)

	return data
}