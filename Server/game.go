package main

import (
	"context"
	"fmt"
	"log"
)

type GameManager struct {
	Action chan Message
}

type Game struct {
	ID               string       `json:"id"`
	CreatorID        string       `json:"creatorId"`
	CreatorName      string       `json:"creatorName"`
	GameName         string       `json:"gameName"`
	Contestants      []Contestant `json:"contestants"`
	Rounds           []Round      `json:"rounds"`
	Round            int          `json:"round"`
	EnableBuzzers    bool         `json:"enableBuzzers"`
	ActiveCategory   int          `json:"activeCategory"`
	ActiveClue       int          `json:"activeClue"`
	ActiveContestant int          `json:"activeContestant"`
}

type Contestant struct {
	Name     string `json:"name"`
	Media    string `json:"media"`
	Score    int    `json:"score"`
	Disabled bool   `json:"disabled"`
}

type Round struct {
	Name       string     `json:"name"`
	Comment    string     `json:"comment"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Clues   []Clue `json:"clues"`
	Media   string `json:"media"`
}

type Clue struct {
	Comment     string `json:"comment"`
	Answered    bool   `json:"answered"`
	Answer      string `json:"answer"`
	Question    string `json:"question"`
	Media       string `json:"media"`
	DailyDouble bool   `json:"dailyDouble"`
}

type ClueSelect struct {
	CategoryNumber int `json:"categoryNumber"`
	ClueNumber     int `json:"clueNumber"`
}

type ContestantSelect struct {
	Contestant int `json:"contestant"`
}

type ClueDetermination struct {
	CategoryNumber    int  `json:"categoryNumber"`
	ClueNumber        int  `json:"clueNumber"`
	Correct           bool `json:"correct"`
	DailyDoubleAmount int  `json:"dailyDoubleAmount"`
}

func (gm *GameManager) Run(ctx context.Context) {

	gm.Action = make(chan Message, 1000)

	/*if err != nil {
		log.Printf("Unable to parse approvals template: %v", err)
	}*/

	// use time.After or similar to check for expired requests
	for {
		select {
		case <-ctx.Done():
			return
		case message := <-gm.Action:
			if (message == Message{}) {
				log.Println("Got empty message. Possible unknown MessageType.")
				continue
			}

			game := gm.getGame(message.GameID)

			switch message.MessageType {
			case "GameState":
				//var payload = message.Payload.(GameID)

				newMessage, _ := makeMessage("GameState", message.GameID, game)
				message.Client.send <- newMessage
			case "GameStateAll":
				//var payload = message.Payload.(GameID)

				newMessage, _ := makeMessage("GameState", message.GameID, game)
				hub.broadcast <- newMessage
			case "JoinGame":
				log.Printf("Joining Game %s", message.GameID)

				hub.joinGame <- JoinInfo{
					GameID: message.GameID,
					Client: message.Client,
				}

				message.MessageType = "GameState"

				gm.Action <- message

			case "SelectClue":
				//var request ClueSelect
				var payload = message.Payload.(ClueSelect)

				gm.selectClue(&game, payload.CategoryNumber, payload.ClueNumber)

			case "ClueDetermination":
				var payload = message.Payload.(ClueDetermination)

				gm.determineClue(&game, payload.Correct, payload.DailyDoubleAmount)

			case "EnableBuzzers":
				gm.enableBuzzers(&game)

			case "ResetBuzzers":
				gm.resetBuzzers(&game)

			case "SelectContestant":
				//var request ClueSelect
				var payload = message.Payload.(ContestantSelect)

				gm.selectContestant(&game, payload.Contestant)
			} // End of Select
		}
	}
}

func (gm *GameManager) getGame(gameID string) Game {
	var game Game

	databaseManager.Bucket.Get("game:"+gameID, &game)

	//fmt.Printf("DB Doc: %+v\n", game)

	return game

}

func (gm *GameManager) selectClue(game *Game, categoryNumber int, clueNumber int) error {

	game.ActiveCategory = categoryNumber
	game.ActiveClue = clueNumber

	_, err := databaseManager.Bucket.Upsert("game:"+game.ID, game, 0)

	if err != nil {
		return fmt.Errorf("could not write game to DB: %v", err)
	}

	message, err := makeMessage("GameStateAll", game.ID, nil)

	gm.Action <- message

	if err != nil {
		return fmt.Errorf("could not make GameState message: %v", err)
	}

	payload := ClueSelect{
		CategoryNumber: categoryNumber,
		ClueNumber:     clueNumber,
	}

	message, err = makeMessage("SelectClue", game.ID, payload)

	if err != nil {
		return fmt.Errorf("could not make SelectClue message: %v", err)
	}

	hub.broadcast <- message

	return nil
}

func (gm *GameManager) enableBuzzers(game *Game) error {

	game.EnableBuzzers = true

	_, err := databaseManager.Bucket.Upsert("game:"+game.ID, game, 0)

	if err != nil {
		return fmt.Errorf("could not write game to DB: %v", err)
	}

	message, err := makeMessage("GameStateAll", game.ID, nil)

	gm.Action <- message

	if err != nil {
		return fmt.Errorf("could not make GameState message: %v", err)
	}

	message, err = makeMessage("EnableBuzzers", game.ID, nil)

	if err != nil {
		return fmt.Errorf("could not make EnableBuzzers message: %v", err)
	}

	hub.broadcast <- message

	return nil
}

func (gm *GameManager) resetBuzzers(game *Game) error {

	game.ActiveContestant = -1
	game.EnableBuzzers = false

	for i := range game.Contestants {
		game.Contestants[i].Disabled = false
	}

	_, err := databaseManager.Bucket.Upsert("game:"+game.ID, game, 0)

	if err != nil {
		return fmt.Errorf("could not write game to DB: %v", err)
	}

	message, err := makeMessage("GameStateAll", game.ID, nil)

	gm.Action <- message

	if err != nil {
		return fmt.Errorf("could not make GameState message: %v", err)
	}

	message, err = makeMessage("ResetBuzzers", game.ID, nil)

	if err != nil {
		return fmt.Errorf("could not make ResetBuzzers message: %v", err)
	}

	hub.broadcast <- message

	return nil
}

func (gm *GameManager) selectContestant(game *Game, contestant int) error {

	game.ActiveContestant = contestant
	game.EnableBuzzers = false

	_, err := databaseManager.Bucket.Upsert("game:"+game.ID, game, 0)

	if err != nil {
		return fmt.Errorf("could not write game to DB: %v", err)
	}

	message, err := makeMessage("GameStateAll", game.ID, nil)

	gm.Action <- message

	if err != nil {
		return fmt.Errorf("could not make GameState message: %v", err)
	}

	message, err = makeMessage("SelectContestant", game.ID, nil)

	if err != nil {
		return fmt.Errorf("could not make SelectContestant message: %v", err)
	}

	hub.broadcast <- message

	return nil
}

func (gm *GameManager) determineClue(game *Game, correct bool, dailyDoubleAmount int) error {
	previousCategory := game.ActiveCategory
	previousClue := game.ActiveClue

	clueValue := (game.Round + game.ActiveClue + 1) * 200

	game.Rounds[game.Round].Categories[game.ActiveCategory].Clues[game.ActiveClue].Answered = true

	if correct {
		for i := range game.Contestants {
			game.Contestants[i].Disabled = false
		}

		game.Contestants[game.ActiveContestant].Score += clueValue

		log.Printf("Addint %v to contestant %v score", clueValue, game.ActiveContestant)

		game.ActiveCategory = -1
		game.ActiveClue = -1
		game.ActiveContestant = -1
		game.EnableBuzzers = false
	}

	if !correct {
		game.Contestants[game.ActiveContestant].Disabled = true

		game.Contestants[game.ActiveContestant].Score -= clueValue

		log.Printf("Deductin %v from contestant %v score", clueValue, game.ActiveContestant)

		game.ActiveContestant = -1
		game.EnableBuzzers = true
	}

	_, err := databaseManager.Bucket.Upsert("game:"+game.ID, game, 0)

	if err != nil {
		return fmt.Errorf("could not write game to DB: %v", err)
	}

	message, err := makeMessage("GameStateAll", game.ID, nil)

	gm.Action <- message

	if err != nil {
		return fmt.Errorf("could not make GameState message: %v", err)
	}

	payload := ClueDetermination{
		CategoryNumber:    previousCategory,
		ClueNumber:        previousClue,
		Correct:           correct,
		DailyDoubleAmount: dailyDoubleAmount,
	}

	message, err = makeMessage("ClueDetermination", game.ID, payload)

	if err != nil {
		return fmt.Errorf("could not make ClueDetermination message: %v", err)
	}

	hub.broadcast <- message

	return nil
}
