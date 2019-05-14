package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Message struct {
	MessageType string      `json:"messageType,omitempty"`
	Timestamp   time.Time   `json:"timestamp,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
	//Payload []byte `json:"payload,omitempty"`
}

func UnmarshallMessage(rawMessage []byte) Message {
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

	case "SelectClue":
		var payload struct {
			Payload ClueSelect `json:"payload,omitempty"`
		}

		err := json.Unmarshal(rawMessage, &payload)

		if err != nil {
			log.Printf("Could not unmarshall message payload: %v\n", err)
			return Message{}
		}

		fmt.Printf("%+v\n", payload)

		return Message{
			MessageType: messageMetatdata.MessageType,
			Timestamp:   messageMetatdata.Timestamp,
			Payload:     payload.Payload,
		}
	}

	return Message{}
}

func makeMessage(MessageType string, Payload interface{}) ([]byte, error) {

	var message Message

	message.MessageType = MessageType
	message.Timestamp = time.Now()
	message.Payload = Payload

	messageJSON := []byte{}
	var err error

	if localSettings.Production {
		messageJSON, err = json.Marshal(message)
	}

	if !localSettings.Production {
		messageJSON, err = json.MarshalIndent(message, "", "    ")
	}

	//fmt.Printf(string(messageJSON))

	if err != nil {
		return nil, fmt.Errorf("Error writing results to JSON: %v", err)
	}

	return messageJSON, nil

}
