package system_health

import (
	"TemplateApi/src/models"
	"TemplateApi/src/service"
	_ "TemplateApi/src/service"
	"time"
)

//type local_Service struct{ //Causes circular dependency?
//	service.Service
//}

//type local_Service service.TemplateService //Still causing circular dependency

type local_service struct {
	service.TemplateService
}

type HealthReporter interface {
	Heartbeat() (*models.HeartbeatResponse, error)
}

func (s local_service) Heartbeat() (*models.HeartbeatResponse, error) {
	return &models.HeartbeatResponse{
		Message: "API Running",
		Time:    time.Now(),
	}, nil
}
