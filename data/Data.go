package data

type Data struct {
	Game  Game   `json:"game"`
	Turn  string `json:"turn"`
	Board Board  `json:"board"`
	You   Snake  `json:"you"`
}

// func ParseData(game, turn, board, you) Data {
// 	data := Data{
// 		Game: Game{
// 			Id: game["id"],
// 			Ruleset: Ruleset{
// 				Name:    game["ruleset"]["name"],
// 				Version: game["ruleset"]["version"],
// 			},
// 			Timeout: game["timeout"],
// 		},
// 		Turn: turn,
// 		Board: Board{
// 			Height: board["height"],
// 			Width:  board["width"],
// 			Food: board["food"],
// 			Hazards: board["hazards"],
// 			Snakes: ,
// 		},
// 	}

// 	return data
// }