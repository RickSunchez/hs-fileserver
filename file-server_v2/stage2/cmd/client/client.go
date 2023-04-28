package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:23456")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Client started!")

	request := "Give me everything you have!"
	_, err = connection.Write([]byte(request))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Sent: %s\n", request)

	inputStream := make([]byte, 1024)
	length, err := connection.Read(inputStream)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Received: %s\n", string(inputStream[:length]))
	connection.Close()
}
