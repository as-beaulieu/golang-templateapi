package service

import (
	"TemplateApi/src/models"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type Messenger interface {
	CreateSimpleMessage(message models.SimpleMessage) (*models.SimpleMessageResponse, error)
}

func (s service) CreateSimpleMessage(message models.SimpleMessage) (*models.SimpleMessageResponse, error) {
	newFile, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("failure creating new file", newFile.Name())
		return nil, err
	}
	defer newFile.Close()

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(message)
	if err != nil {
		fmt.Println("failure encoding message to buffer")
		return nil, err
	}

	_, err = buffer.WriteTo(newFile)
	if err != nil {

	}

	return &models.SimpleMessageResponse{
		ID:      message.ID,
		Message: message.Message,
		Error:   nil,
	}, nil
}
