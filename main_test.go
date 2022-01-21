package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Tch1b0/Esproso/brain"
	GameData "github.com/Tch1b0/Esproso/data"
)

type SampleRequest struct {
	Data             GameData.Data `json:"data"`
	ExpectedResponse string        `json:"expect"`
}

func TestBrain(t *testing.T) {
	text, _ := ioutil.ReadFile("./samples/sample_requests.json")
	samples := []SampleRequest{}
	if err := json.Unmarshal(text, &samples); err != nil {
		fmt.Print(err)
		return
	}

	for _, sample := range samples {
		t.Run("Test Move", func(t *testing.T) {
			if ans := brain.Move(sample.Data); ans != sample.ExpectedResponse {
				t.Errorf("Expected move \"%s\", but received \"%s\"", sample.ExpectedResponse, ans)
			}
		})
	}
}
