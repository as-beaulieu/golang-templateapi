package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // here
	"log"
	"os"
	"strconv"
)

// TODO Now need to read from .sql files
// TODO iterate through multiple .sql files so can keep a repository of sql files

func main() {
	err := godotenv.Load()
	psqlInfo, err := getPsqlInfo()
	if err != nil {
		log.Fatal("problem getting psqlInfo: ", err)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to open DB connection: ", err)
	}
	defer db.Close()

	createSql := `CREATE TABLE IF NOT EXISTS public.users (
id TEXT primary key,
name TEXT
) ;`

	_, err = db.Exec(createSql)
	if err != nil {
		log.Fatal("error executing sql script: ", err)
	}

	fmt.Println("complete")
}

func getPsqlInfo() (string, error) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		return "", fmt.Errorf("host is empty from env variables")
	}

	var port int
	var err error
	portString := os.Getenv("POSTGRES_PORT")
	if portString == "" {
		return "", fmt.Errorf("port is empty from env variables")
	} else {
		port, err = strconv.Atoi(portString)
		if err != nil {
			return "", fmt.Errorf("unable to convert port env from string to int")
		}
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		return "", fmt.Errorf("user is empty from env variables")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		return "", fmt.Errorf("password is empty from env variables")
	}

	dbname := os.Getenv("POSTGRES_DATABASE_NAME")
	if dbname == "" {
		return "", fmt.Errorf("dbname is empty from env variables")
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password), nil
}
