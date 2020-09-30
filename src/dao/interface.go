package dao

import (
	"TemplateApi/src/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //remember to import the concrete implementation of the db driver, and initialize
)

type DAO interface {
	CreateUser(user models.User) error
	GetUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id string) error
}

type dao struct {
	host       string
	port       int
	user       string
	password   string
	dbname     string
	connection *sql.DB
}

type PostgresBuilder struct {
	dao
}

func (pb PostgresBuilder) SetHost(host string) PostgresBuilder {
	a := pb
	a.host = host
	return a
}

func (pb PostgresBuilder) SetPort(port int) PostgresBuilder {
	a := pb
	a.port = port
	return a
}

func (pb PostgresBuilder) SetUser(user string) PostgresBuilder {
	a := pb
	a.user = user
	return a
}

func (pb PostgresBuilder) SetPassword(pass string) PostgresBuilder {
	a := pb
	a.password = pass
	return a
}

func (pb PostgresBuilder) SetDbName(name string) PostgresBuilder {
	a := pb
	a.dbname = name
	return a
}

func (pb PostgresBuilder) Build() *dao {
	var err error
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	pb.host, pb.port, pb.user, pb.password, pb.dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		pb.host, pb.port, pb.user, pb.password)

	pb.connection, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err) //TODO: either import logger into DAO as well, or do something better than panic!
	}

	return &pb.dao
}
