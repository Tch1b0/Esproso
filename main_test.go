package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Tch1b0/Esproso/brain"
	GameData "github.com/Tch1b0/Esproso/data"
)

func TestMoves(t *testing.T) {
	text, _ := ioutil.ReadFile("./samples/sample_requests.json")
	data := []GameData.Data{}
	if err := json.Unmarshal(text, &data); err != nil {
		fmt.Print(err)
		return
	}

	t.Run("Test Move", func(t *testing.T) {
		expected := "up"
		if ans := brain.Move(data[0]); ans != "up" {
			t.Errorf("Expected move \"%s\", but received \"%s\"", expected, ans)
		}
	})
}
