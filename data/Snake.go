package data

type Snake struct {
	Id      string       `json:"id"`
	Name    string       `json:"name"`
	Health  int          `json:"health"`
	Body    []Coordinate `json:"body"`
	Latency string       `json:"latency"`
	Head    Coordinate   `json:"head"`
	Length  int          `json:"length"`
	Shout   string       `json:"shout"`
	Squad   string       `json:"squad"`
}

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}
