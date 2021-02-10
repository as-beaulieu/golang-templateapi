package service

import (
	"TemplateApi/src/dao"
	"TemplateApi/src/service/message"
	"TemplateApi/src/service/system_health"
	"TemplateApi/src/service/user"
	"TemplateApi/src/service/weather"
	"go.uber.org/zap"
)

type Service interface {
	//PublicFunctionName(input) (output, error)
	system_health.HealthReporter
	message.Messenger
	user.UserOperator
	weather.WeatherReporter
}

type TemplateService struct {
	//package	PackageType
	Logger   *zap.Logger
	Postgres *dao.DAO
}

type ServiceBuilder struct {
	TemplateService
}

//func (sb ServiceBuilder) WithDIPackage(package PackageType) ServiceBuilder {
//	a := sb
//	a.package = package
//	return a
//}

func (sb *ServiceBuilder) WithLogger(logger *zap.Logger) ServiceBuilder {
	a := *sb
	a.Logger = logger
	return a
}

func (sb ServiceBuilder) WithPostgres(dao dao.DAO) ServiceBuilder { //Point to Interface of package to be injected
	a := sb
	a.Postgres = dao
	return a
}

func (sb ServiceBuilder) Build() *TemplateService {
	return &sb.TemplateService
}
