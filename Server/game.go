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

type GameID struct {
	GameID string `json:"gameId"`
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

			switch message.MessageType {
			case "GameState":
				//var payload = message.Payload.(GameID)

				newMessage, _ := makeMessage("GameState", `{"id":"1234","categories":[{"name":"RECENT BESTSELLERS","clues":[{"answer":"\"Fire & Blood\" is a novel from this author, set centuries before the events in \"Game of Thrones\"","question":"(George) Martin","dailyDouble":false,"answered":false},{"answer":"Tara Westover wrote of her quest to be this, the 1-word title of her memoir about leaving her survivalist family to go to college","question":"Educated","dailyDouble":false,"answered":false},{"answer":"Both Michael Wolff's \"Fire and Fury\" & Bob Woodward's \"Fear\" mention this building in their subtitles","question":"the White House","dailyDouble":true,"answered":false},{"answer":"This bard of Bangor revisited the bestseller list with the creepy entry \"The Outsider\"","question":"Stephen King","dailyDouble":false,"answered":false},{"answer":"Marlon Bundo is a type of this animal in books from Mike Pence's daughter & from John Oliver's \"Last Week Tonight\"","question":"a rabbit","dailyDouble":false,"answered":false}]},{"name":"WOMEN OLYMPIANS","clues":[{"answer":"American saber specialist Ibtihaj Muhammad wore a hijab competing in this Olympic sport in 2016","question":"fencing","dailyDouble":false,"answered":false},{"answer":"In 2016 tennis ace Monica Puig beat 3 Grand Slam champions to win this U.S. territory's first-ever gold medal","question":"Puerto Rico","dailyDouble":false,"answered":false},{"answer":"China's Li Xuerui took gold in 2012 in this racquet sport","question":"badminton","dailyDouble":false,"answered":false},{"answer":"Jackie Joyner-Kersee twice won gold in this multi-event Olympic contest","question":"the heptathlon","dailyDouble":false,"answered":false},{"answer":"Simone Biles won 4 gold medals in Rio: team, all-around, floor exercise & this high-flying event","question":"the vault","dailyDouble":false,"answered":false}]},{"name":"THE ELEMENTS","clues":[{"answer":"Alcoa is the nation's largest producer of this element","question":"aluminum","dailyDouble":false,"answered":false},{"answer":"Boy, if I had one of this element, atomic No. 28, every time someone said it was first isolated in 1751","question":"nickel","dailyDouble":false,"answered":false},{"answer":"Protactinium was once called brevium because the first isotope discovered had only a 70-second one of these","question":"a half-life","dailyDouble":false,"answered":false},{"answer":"The first 3 halogen or \"salt-forming\" elements on the periodic table are fluorine, this, bromine","question":"chlorine","dailyDouble":false,"answered":false},{"answer":"This element precedes carbonate in a compound used to treat gastric ulcers & is definitely some of your...","question":"bismuth (or calcium)","dailyDouble":false,"answered":false}]},{"name":"SPELL THE BEANS","comment":"Each response is a type of bean that you're going to have to spell for us.","clues":[{"answer":"Its name comes from being shaped like an internal organ","question":"K-I-D-N-E-Y","dailyDouble":false,"answered":false},{"answer":"A chocolate source--not the spelling Swiss Miss uses","question":"C-A-C-A-O","dailyDouble":false,"answered":false},{"answer":"Sizzling cinnamon & juicy pear are 2 popular flavors of it","question":"J-E-L-L-Y","dailyDouble":false,"answered":false},{"answer":"Geographic AKA of butter beans","question":"L-I-M-A","dailyDouble":false,"answered":false},{"answer":"Chipotle's ingredients include black beans & this other 5-letter type","question":"P-I-N-T-O","dailyDouble":false,"answered":false}]},{"name":"WORDS IN TEACHER","clues":[{"answer":"Third from the Sun","question":"Earth","dailyDouble":false,"answered":false},{"answer":"To provide food for an event","question":"cater","dailyDouble":false,"answered":false},{"answer":"To deceive for profit","question":"cheat","dailyDouble":false,"answered":false},{"answer":"It was originally reckoned as the amount of land a team of oxen could plow in a day","question":"an acre","dailyDouble":false,"answered":false},{"answer":"Jasper & agate are 2 of the many minerals classified as this form of quartz","question":"chert","dailyDouble":false,"answered":false}]},{"name":"EDUCATIONAL GAMES","clues":[{"answer":"This \"amphibious\" company makes the educational line of LeapPad tablets","question":"LeapFrog","dailyDouble":false,"answered":false},{"answer":"In a geography computer game, you need to find out \"Where in the World Is\" she","question":"Carmen Sandiego","dailyDouble":false,"answered":false},{"answer":"The Amazon Trail is a spin-off from this classic history game","question":"The Oregon Trail","dailyDouble":false,"answered":false},{"answer":"Ms. Frizzle & the yellow title transport feature in web games based on this TV show","question":"The Magic School Bus","dailyDouble":false,"answered":false},{"answer":"Mavis Beacon, who has taught this life skill to millions, is not a real person","question":"typing","dailyDouble":false,"answered":false}]}]}`)

				message.Client.send <- newMessage
			case "JoinGame":
				var payload = message.Payload.(GameID)

				log.Printf("Joining Game %s", payload.GameID)

				hub.joinGame <- JoinInfo{
					GameID: payload.GameID,
					Client: message.Client,
				}

				message.MessageType = "GameState"

				gm.Action <- message

			case "SelectClue":
				//var request ClueSelect
				var payload = message.Payload.(ClueSelect)

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
