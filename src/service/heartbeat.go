package service

import (
	"TemplateApi/src/models"
	"time"
)

func (s service) Heartbeat() (*models.HeartbeatResponse, error) {
	return &models.HeartbeatResponse{
		Message: "API Running",
		Time:    time.Now(),
	}, nil
}
