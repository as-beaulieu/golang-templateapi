package message

import (
	"TemplateApi/src/models"
	"TemplateApi/src/service"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type local_service struct {
	*service.TemplateService
}

type Messenger interface {
	CreateSimpleMessage(message models.SimpleMessage) (*models.SimpleMessageResponse, error)
}

func (s local_service) CreateSimpleMessage(message models.SimpleMessage) (*models.SimpleMessageResponse, error) {
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
