package brain

import (
	"fmt"

	GameData "github.com/Tch1b0/Esproso/data"
)


func Move(data GameData.Data) string {
	fmt.Printf("Here")
	position := data.You.Head
	possibleMoves := []string{
		"up",
		"down",
		"left",
		"right",
	}
	avoidMoves := avoid(data)

	closeFood := findFood(data)

	// Get the direction the player has to go to get to the food
	foodDirection := getDirection(position, closeFood)

	// Go up/down if y-value isn't equal
	if position.Y != closeFood.Y && 
	    !containsString(avoidMoves, foodDirection) {
		return foodDirection
	} else if position.X != closeFood.X && 
		!containsString(avoidMoves, foodDirection) {
		return foodDirection		
	}

	for _, move := range possibleMoves {
		if !containsString(avoidMoves, move) {
			return move
		}
	}

	fmt.Printf("No escape")
	return "up"
}