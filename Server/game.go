package main

import (
	"context"
	"log"
)

type GameManager struct {
	Action chan Message
}

type Game struct {
	ID             string       `json:"id"`
	CreatorID      string       `json:"creatorId"`
	CreatorName    string       `json:"creatorName"`
	GameName       string       `json:"gameName"`
	Contestants    []Contestant `json:"players"`
	Rounds         []Round      `json:"rounds"`
	Round          int          `json:"round"`
	ActiveCategory Category     `json:"activeCategory"`
	ActiveClue     Clue         `json:"activeClue"`
}

type Contestant struct {
	Name  string `json:"name"`
	Media string `json:"media"`
	Score string `json:"score"`
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
	Round          int `json:"round"`
	CategoryNumber int `json:"categoryNumber"`
	ClueNumber     int `json:"clueNumber"`
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

				newMessage, _ := makeMessage("GameState", message.GameID, `{"id":"a267dd0b-40cb-4178-ad8c-58d5efa1ff29","round":0,"rounds":[{"name":"Round 1","categories":[{"name":"RECENT BESTSELLERS","clues":[{"question":"(George) Martin","answer":"\"Fire & Blood\" is a novel from this author, set centuries before the events in \"Game of Thrones\"","dailyDouble":false,"answered":false},{"question":"Educated","answer":"Tara Westover wrote of her quest to be this, the 1-word title of her memoir about leaving her survivalist family to go to college","dailyDouble":false,"answered":false},{"question":"the White House","answer":"Both Michael Wolff's \"Fire and Fury\" & Bob Woodward's \"Fear\" mention this building in their subtitles","dailyDouble":true,"answered":false},{"question":"Stephen King","answer":"This bard of Bangor revisited the bestseller list with the creepy entry \"The Outsider\"","dailyDouble":false,"answered":false},{"question":"a rabbit","answer":"Marlon Bundo is a type of this animal in books from Mike Pence's daughter & from John Oliver's \"Last Week Tonight\"","dailyDouble":false,"answered":false}]},{"name":"WOMEN OLYMPIANS","clues":[{"question":"fencing","answer":"American saber specialist Ibtihaj Muhammad wore a hijab competing in this Olympic sport in 2016","dailyDouble":false,"answered":false},{"question":"Puerto Rico","answer":"In 2016 tennis ace Monica Puig beat 3 Grand Slam champions to win this U.S. territory's first-ever gold medal","dailyDouble":false,"answered":false},{"question":"badminton","answer":"China's Li Xuerui took gold in 2012 in this racquet sport","dailyDouble":false,"answered":false},{"question":"the heptathlon","answer":"Jackie Joyner-Kersee twice won gold in this multi-event Olympic contest","dailyDouble":false,"answered":false},{"question":"the vault","answer":"Simone Biles won 4 gold medals in Rio: team, all-around, floor exercise & this high-flying event","dailyDouble":false,"answered":false}]},{"name":"THE ELEMENTS","clues":[{"question":"aluminum","answer":"Alcoa is the nation's largest producer of this element","dailyDouble":false,"answered":false},{"question":"nickel","answer":"Boy, if I had one of this element, atomic No. 28, every time someone said it was first isolated in 1751","dailyDouble":false,"answered":false},{"question":"a half-life","answer":"Protactinium was once called brevium because the first isotope discovered had only a 70-second one of these","dailyDouble":false,"answered":false},{"question":"chlorine","answer":"The first 3 halogen or \"salt-forming\" elements on the periodic table are fluorine, this, bromine","dailyDouble":false,"answered":false},{"question":"bismuth (or calcium)","answer":"This element precedes carbonate in a compound used to treat gastric ulcers & is definitely some of your...","dailyDouble":false,"answered":false}]},{"name":"SPELL THE BEANS","comment":"Each response is a type of bean that you're going to have to spell for us.","clues":[{"question":"K-I-D-N-E-Y","answer":"Its name comes from being shaped like an internal organ","dailyDouble":false,"answered":false},{"question":"C-A-C-A-O","answer":"A chocolate source--not the spelling Swiss Miss uses","dailyDouble":false,"answered":false},{"question":"J-E-L-L-Y","answer":"Sizzling cinnamon & juicy pear are 2 popular flavors of it","dailyDouble":false,"answered":false},{"question":"L-I-M-A","answer":"Geographic AKA of butter beans","dailyDouble":false,"answered":false},{"question":"P-I-N-T-O","answer":"Chipotle's ingredients include black beans & this other 5-letter type","dailyDouble":false,"answered":false}]},{"name":"WORDS IN TEACHER","clues":[{"question":"Earth","answer":"Third from the Sun","dailyDouble":false,"answered":false},{"question":"cater","answer":"To provide food for an event","dailyDouble":false,"answered":false},{"question":"cheat","answer":"To deceive for profit","dailyDouble":false,"answered":false},{"question":"an acre","answer":"It was originally reckoned as the amount of land a team of oxen could plow in a day","dailyDouble":false,"answered":false},{"question":"chert","answer":"Jasper & agate are 2 of the many minerals classified as this form of quartz","dailyDouble":false,"answered":false}]},{"name":"EDUCATIONAL GAMES","clues":[{"question":"LeapFrog","answer":"This \"amphibious\" company makes the educational line of LeapPad tablets","dailyDouble":false,"answered":false},{"question":"Carmen Sandiego","answer":"In a geography computer game, you need to find out \"Where in the World Is\" she","dailyDouble":false,"answered":false},{"question":"The Oregon Trail","answer":"The Amazon Trail is a spin-off from this classic history game","dailyDouble":false,"answered":false},{"question":"The Magic School Bus","answer":"Ms. Frizzle & the yellow title transport feature in web games based on this TV show","dailyDouble":false,"answered":false},{"question":"typing","answer":"Mavis Beacon, who has taught this life skill to millions, is not a real person","dailyDouble":false,"answered":false}]}]}]}`)

				message.Client.send <- newMessage
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

				gm.selectClue(message.GameID, payload.Round, payload.CategoryNumber, payload.ClueNumber)
			}
		}
	}
}

func (gm *GameManager) selectClue(gameID string, round int, categoryNumber int, clueNumber int) {
	payload := ClueSelect{
		Round:          round,
		CategoryNumber: categoryNumber,
		ClueNumber:     clueNumber,
	}

	message, err := makeMessage("SelectClue", gameID, payload)

	if err != nil {
		log.Printf("Unable to make message: %v", err)
		return
	}

	hub.broadcast <- message
}
