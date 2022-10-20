package data

type Board struct {
	Height  int          `json:"height"`
	Width   int          `json:"width"`
	Food    []Coordinate `json:"food"`
	Hazards []Coordinate `json:"hazards"`
	Snakes  []Snake      `json:"snakes"`
}
