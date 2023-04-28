package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:23456")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server started!")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	inputStream := make([]byte, 1024)
	length, err := connection.Read(inputStream)
	if err != nil {
		connection.Close()
		return
	}

	request := string(inputStream[:length])
	fmt.Printf("Received: %s\n", request)
	if request != "Give me everything you have!" {
		connection.Close()
		return
	}

	response := "All files were sent!"
	connection.Write([]byte(response))

	fmt.Printf("Sent: %s\n", response)
	connection.Close()
	os.Exit(0)
}
