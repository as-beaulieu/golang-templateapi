package main

import (
	"TemplateApi/src/dao"
	"TemplateApi/src/models"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
)

var (
	users     = make([]models.User, 0)
	userNames = []string{
		"bob",
		"mary",
		"john",
		"jane",
		"bill",
		"bo",
	}
)

type Configuration struct {
	DbPort string
	DbName string
	DbUser string
	DbPass string
	DbHost string
}

func main() {
	config := &Configuration{}
	flag.StringVar(&config.DbPort, "DATABASE_PORT", os.Getenv("POSTGRES_PORT"), "the environment variable for database port")
	flag.StringVar(&config.DbUser, "DATABASE_USER", os.Getenv("POSTGRES_USER"), "the username for the database connection")
	flag.StringVar(&config.DbPass, "DATABASE_PASS", os.Getenv("POSTGRES_PASSWORD"), "the password for the database connection")
	flag.StringVar(&config.DbName, "DATABASE_NAME", os.Getenv("POSTGRES_DATABASE_NAME"), "the name of the database")
	flag.StringVar(&config.DbHost, "DATABASE_HOST", os.Getenv("POSTGRES_HOST"), "the host root for connecting to the database")
	flag.Parse()

	fmt.Println(&config)

	port, err := strconv.Atoi(config.DbPort)
	if err != nil {
		fmt.Println("error converting port to integer")
		os.Exit(1)
	}

	postgres := dao.PostgresBuilder{}.
		SetPort(port).
		SetHost(config.DbHost).
		SetPassword(config.DbPass).
		SetUser(config.DbUser).
		Build()

	users := generateUsers()

	for _, user := range users {
		err = postgres.CreateUser(user)
		if err != nil {
			fmt.Println("error saving user to database", err)
			os.Exit(1)
		}
	}

	fmt.Println("completed")
}

func generateUsers() []models.User {
	for _, name := range userNames {
		user := models.User{
			ID:   uuid.New().String(),
			Name: name,
		}
		users = append(users, user)
	}
	return users
}
