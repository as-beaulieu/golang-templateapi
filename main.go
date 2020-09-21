package main

import (
	"TemplateApi/src/dao"
	"TemplateApi/src/logging"
	"TemplateApi/src/server"
	"TemplateApi/src/service"
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	errors      chan error
	tcpResponse chan string
)

func main() {
	err := godotenv.Load()
	example := os.Getenv("EXAMPLE")
	fmt.Println(example)

	errors = make(chan error)
	tcpResponse = make(chan string)

	dbport := os.Getenv("POSTGRES_PORT")
	port, err := strconv.Atoi(dbport)

	logger := logging.NewLogger()
	postgres := dao.PostgresBuilder{}.
		SetUser(os.Getenv("POSTGRES_USER")).
		SetPassword(os.Getenv("POSTGRES_PASSWORD")).
		SetHost(os.Getenv("POSTGRES_HOST")).
		SetPort(port).
		SetDbName(os.Getenv("POSTGRES_HOST")).
		Build()
	svc := service.ServiceBuilder{}.
		WithLogger(*logger).
		WithPostgres(postgres).
		Build()

	//http server
	go func() {
		errors <- server.RunHttpServer(svc)
	}()

	//grpc server

	//tcp server
	go func() {
		tcpServer, err := net.Listen("tcp", ":9090")
		if err != nil {
			errors <- err
		}

		for {
			conn, err := tcpServer.Accept()
			if err != nil {
				errors <- err
			}
			go handleConn(conn)
		}
	}()

	err = <-errors
	close(errors)
	if err != nil {
		log.Fatal("error: ", err)
	}
	log.Println("program closed")
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "enter a new number:")

	scanner := bufio.NewScanner(conn)

	//take in the number from stdin and return it with an acknowledgement
	go func() {
		for scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			response := "received number: " + string(rune(num))
			tcpResponse <- response
			io.WriteString(conn, "\nEnter a new number:")
		}
	}()

	for _ = range tcpResponse {
		fmt.Println(tcpResponse)
	}
}
