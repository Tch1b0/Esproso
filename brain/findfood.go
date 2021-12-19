package brain

import (
	"math"

	GameData "github.com/Tch1b0/Esproso/data"
)

func findFood(data GameData.Data) GameData.Coordinate {
	position := data.You.Head
	var lowestDistance float64 = -1
	var closestFood GameData.Coordinate

	for _, food := range data.Board.Food {
		distY := position.Y - food.Y
		distX := position.X - food.X
		distance := math.Sqrt(math.Pow(float64(distY), 2) * math.Pow(float64(distX), 2))

		if distance < lowestDistance || lowestDistance == -1 {
			lowestDistance = distance
			closestFood = food
		}
	}

	return closestFood
}
