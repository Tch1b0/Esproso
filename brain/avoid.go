package brain

import (
	GameData "github.com/Tch1b0/Esproso/data"
)


func avoid(data GameData.Data) []string {
	var avoidMoves []string
	position := data.You.Head
	surroundings := []*GameData.Coordinate{
		{
			X: position.X + 1, 
			Y: position.Y,
		},
		{
			X: position.X -1,
			Y: position.Y, 
		},
		{
			X: position.X,
			Y: position.Y + 1,
		},
		{
			X: position.X,
			Y: position.Y - 1,
		},
	}

	// Check if the snake is at the border of the y axis
	if position.Y + 1 == data.Board.Height {
		avoidMoves = append(avoidMoves, "up")
	} else if position.Y == 0 {
		avoidMoves = append(avoidMoves, "down")
	}

	// Check if the snake is at the border of the x axis
	if position.X + 1 == data.Board.Width {
		avoidMoves = append(avoidMoves, "right")
	} else if position.X == 0 {
		avoidMoves = append(avoidMoves, "left")
	}

	for _, pos := range data.You.Body {
		if containsCoordinate(surroundings, pos) {
			avoidMoves = append(avoidMoves, getDirection(position, pos))
		}
	}

	for _, snake := range data.Board.Snakes {
		if containsCoordinate(surroundings, snake.Head) {
			avoidMoves = append(avoidMoves, getDirection(position, snake.Head))
		}
		for _, pos := range snake.Body {
			if containsCoordinate(surroundings, pos) {
				avoidMoves = append(avoidMoves, getDirection(position, pos))
			}
		}
	}

	return avoidMoves
}

