package dao

import (
	"TemplateApi/src/models"
	"database/sql"
	"fmt"
)

type DAO interface {
	CreateUser(user models.User) error
	GetUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id string) error
}

type dao struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	db       *sql.DB
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pb.host, pb.port, pb.user, pb.password, pb.dbname)
	pb.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer pb.db.Close()

	return &pb.dao
}
