package main

import (
	"encoding/json"
	"log"
	"time"
)

type Message struct {
	Client      *Client
	MessageType string      `json:"messageType,omitempty"`
	Timestamp   time.Time   `json:"timestamp,omitempty"`
	GameID      string      `json:"gameId,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
	//Payload []byte `json:"payload,omitempty"`
}

func UnmarshallMessage(client *Client, rawMessage []byte) Message {

	var messageMetatdata struct {
		MessageType string    `json:"messageType,omitempty"`
		Timestamp   time.Time `json:"timestamp,omitempty"`
	}

	err := json.Unmarshal(rawMessage, &messageMetatdata)

	if err != nil {
		log.Printf("Unable to unmarshall message wrapper: %v", err)
		return Message{}
	}

	switch messageMetatdata.MessageType {
	case "GameState":

	case "JoinGame":
		var payload struct {
			Payload GameID `json:"payload,omitempty"`
		}

		err := json.Unmarshal(rawMessage, &payload)

		if err != nil {
			log.Printf("Could not unmarshall message payload: %v\n", err)
			return Message{}
		}

		newMessage := Message{
			Client:      client,
			MessageType: messageMetatdata.MessageType,
			Timestamp:   messageMetatdata.Timestamp,
			Payload:     payload.Payload,
		}

		log.Printf("Unmarshalled message: %+v", newMessage)

		return newMessage

	case "SelectClue":
		var payload struct {
			Payload ClueSelect `json:"payload,omitempty"`
		}

		/*err := json.Unmarshal(rawMessage, &payload)

		if err != nil {
			log.Printf("Could not unmarshall message payload: %v\n", err)
			return Message{}
		}

		fmt.Printf("%+v\n", payload)*/

		newMessage := Message{
			Client:      client,
			MessageType: messageMetatdata.MessageType,
			Timestamp:   messageMetatdata.Timestamp,
			Payload:     payload.Payload,
		}

		log.Panicf("Unmarshalled message: %+v", newMessage)

		return newMessage
	}

	return Message{}
}

func makeMessage(MessageType string, Payload interface{}) (Message, error) {

	var message Message

	message.MessageType = MessageType
	message.Timestamp = time.Now()
	message.Payload = Payload

	return message, nil

}
