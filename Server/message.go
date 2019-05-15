package main

import (
	"encoding/json"
	"log"
	"time"
)

type Message struct {
	Client      *Client     `json:"client,omitempty"`
	MessageType string      `json:"messageType,omitempty"`
	GameID      string      `json:"gameId,omitempty"`
	Timestamp   time.Time   `json:"timestamp,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
	//Payload []byte `json:"payload,omitempty"`
}

func UnmarshallMessage(client *Client, rawMessage []byte) Message {

	var messageMetatdata struct {
		MessageType string    `json:"messageType,omitempty"`
		GameID      string    `json:"gameId,omitempty"`
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

		newMessage := Message{
			Client:      client,
			GameID:      messageMetatdata.GameID,
			MessageType: messageMetatdata.MessageType,
			Timestamp:   messageMetatdata.Timestamp,
			Payload:     nil,
		}

		//log.Printf("Unmarshalled message: %+v", newMessage)

		return newMessage

	case "SelectClue":
		var payload struct {
			Payload ClueSelect `json:"payload,omitempty"`
		}

		err := json.Unmarshal(rawMessage, &payload)

		if err != nil {
			log.Printf("Could not unmarshall message payload: %v\n", err)
			return Message{}
		}

		//fmt.Printf("%+v\n", payload)

		newMessage := Message{
			GameID:      messageMetatdata.GameID,
			Client:      client,
			MessageType: messageMetatdata.MessageType,
			Timestamp:   messageMetatdata.Timestamp,
			Payload:     payload.Payload,
		}

		log.Printf("Unmarshalled message: %+v\n", newMessage)

		return newMessage
	}

	return Message{}
}

func makeMessage(MessageType string, GameID string, Payload interface{}) (Message, error) {

	var message Message

	message.MessageType = MessageType
	message.GameID = GameID
	message.Timestamp = time.Now()
	message.Payload = Payload

	return message, nil

}
