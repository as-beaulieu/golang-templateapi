package service

import "TemplateApi/src/models"

type Service interface {
	//PublicFunctionName(input) (output, error)
	Heartbeat() (*models.HeartbeatResponse, error)
	GetWeather() (*models.WeatherResponse, error)
	CreateSimpleMessage(message models.SimpleMessage) (*models.SimpleMessageResponse, error)
}

type service struct {
	//package	PackageType
}

type ServiceBuilder struct {
	service
}

//func (sb ServiceBuilder) WithDIPackage(package PackageType) ServiceBuilder {
//	a := sb
//	a.package = package
//	return a
//}

func (sb ServiceBuilder) Build() *service {
	return &sb.service
}