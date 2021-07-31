package data

type Data struct {
	Game  Game   `json:"game"`
	Turn  string `json:"turn"`
	Board Board  `json:"board"`
	You   Snake  `json:"you"`
}
