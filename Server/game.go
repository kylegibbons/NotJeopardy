package main

import (
	"context"
	"log"
)

type GameManager struct {
	Action chan Message
}

type Game struct {
	ID          string     `json:"id,omitempty"`
	CreatorID   string     `json:"creatorId,omitempty"`
	CreatorName string     `json:"creatorName,omitempty"`
	GameName    string     `json:"gameName,omitempty"`
	Round       int        `json:"round,omitempty"`
	Categories  []Category `json:"categories,omitempty"`
}

type Category struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
	Clues   []Clue `json:"clues,omitempty"`
	Media   string `json:"media,omitempty"`
}

type Clue struct {
	Comment     string `json:"comment,omitempty"`
	Answered    bool   `json:"answered,omitempty"`
	Answer      string `json:"answer,omitempty"`
	Question    string `json:"question,omitempty"`
	Media       string `json:"media,omitempty"`
	DailyDouble bool   `json:"dailyDouble,omitempty"`
}

type ClueSelect struct {
	GameID         string `json:"gameId"`
	Round          int    `json:"round"`
	CategoryNumber int    `json:"categoryNumber"`
	ClueNumber     int    `json:"clueNumber"`
}

func (gm *GameManager) Run(ctx context.Context) {

	gm.Action = make(chan Message, 10)

	/*if err != nil {
		log.Printf("Unable to parse approvals template: %v", err)
	}*/

	// use time.After or similar to check for expired requests
	for {
		select {
		case <-ctx.Done():
			return
		case message := <-gm.Action:
			switch message.MessageType {
			case "GameState":
				// return game state
			case "SelectClue":
				//var request ClueSelect
				var payload ClueSelect = message.Payload.(ClueSelect)

				gm.selectClue(payload.GameID, payload.Round, payload.CategoryNumber, payload.ClueNumber)
			}
		}
	}
}

func (gm *GameManager) selectClue(gameID string, round int, categoryNumber int, clueNumber int) {
	payload := ClueSelect{
		GameID:         gameID,
		Round:          round,
		CategoryNumber: categoryNumber,
		ClueNumber:     clueNumber,
	}

	message, err := makeMessage("SelectClue", payload)

	if err != nil {
		log.Printf("Unable to make message: %v", err)
		return
	}

	hub.broadcast <- message
}
