package service

import (
	"TemplateApi/src/dao"
	"go.uber.org/zap"
)

type Service interface {
	//PublicFunctionName(input) (output, error)
	HealthReporter
	Messenger
	UserOperator
	WeatherReporter
}

type service struct {
	//package	PackageType
	logger   zap.Logger
	postgres dao.DAO
}

type ServiceBuilder struct {
	service
}

//func (sb ServiceBuilder) WithDIPackage(package PackageType) ServiceBuilder {
//	a := sb
//	a.package = package
//	return a
//}

func (sb ServiceBuilder) WithLogger(logger zap.Logger) ServiceBuilder {
	a := sb
	a.logger = logger
	return a
}

func (sb ServiceBuilder) WithPostgres(dao dao.DAO) ServiceBuilder { //Point to Interface of package to be injected
	a := sb
	a.postgres = dao
	return a
}

func (sb ServiceBuilder) Build() *service {
	return &sb.service
}
