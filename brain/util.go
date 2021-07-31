package brain

import (
	"fmt"

	GameData "github.com/Tch1b0/Esproso/data")


func containsString(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

func containsCoordinate(arr []*GameData.Coordinate, val GameData.Coordinate) bool {
	for _, a := range arr {
		if *a == val {
			return true
		}
	}
	return false
}

func getDirection(pos1 GameData.Coordinate, pos2 GameData.Coordinate) string {
	fmt.Println("GET DIRECTION")
	if pos1.Y < pos2.Y {
		return "up"
	}
	if pos1.Y > pos2.Y {
		return "down"
	}
	if pos1.X < pos2.X {
		return "right"
	}
	if pos1.X > pos2.X {
		return "left"
	}
	return "Equal"
}