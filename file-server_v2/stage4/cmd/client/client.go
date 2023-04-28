package main

import (
	"bufio"
	clientsocket "file_server/internal/clientSocket"
	"fmt"
	"log"
	"os"
)

const (
	address = "127.0.0.1:23456"
)

var scanner = bufio.NewScanner(os.Stdin)
var helloMessage = "Enter action\n  1 - get a file, 2 - send a file, 3 - delete a file, 4 - list files, 5 - exit:\n"

func main() {
	var input string
	client := clientsocket.New(address)

	err := client.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Client started!")

	for {
		input = cursor(helloMessage)

		switch input {
		case "1": // get
			handleGet(&client)
		case "2": // send
			handleSend(&client)
		case "3": // delete
			handleDelete(&client)
		case "4": // list
			handleList(&client)
		case "5": // exit
			os.Exit(0)
		default:
			fmt.Println("Undefined command")
			continue
		}
	}

}

func handleGet(client *clientsocket.ClientSocket) {
	requestType := cursor("Do you want to get the file by name or by id (1 - name, 2 - id):\n")

	var continueMessage string
	switch requestType {
	case "1":
		continueMessage = "Enter filename:\n"
	case "2":
		continueMessage = "Enter file ID:\n"
	default:
		return
	}

	requestContent := cursor(continueMessage)

	err := client.GetRequest(requestContent, requestType)
	if err != nil {
		fmt.Printf("File %s unavailable\n", requestContent)
		return
	}

	fmt.Printf("File %s loaded:\n", requestContent)
}

func handleSend(client *clientsocket.ClientSocket) {
	cursor("Enter filename:\n")
	filename := scanner.Text()

	fileId, err := client.SendRequest(filename)
	if err != nil {
		fmt.Printf("File %s not sended\n", filename)
		return
	}

	fmt.Printf("The file %s added successfully\n", filename)
	fmt.Printf("File id is \"%d\"\n", fileId)
}

func handleDelete(client *clientsocket.ClientSocket) {
	requestType := cursor("Do you want to delete the file by name or by id (1 - name, 2 - id):\n")

	var continueMessage string
	switch requestType {
	case "1":
		continueMessage = "Enter filename:\n"
	case "2":
		continueMessage = "Enter file ID:\n"
	default:
		return
	}

	requestContent := cursor(continueMessage)

	ok, err := client.DeleteRequest(requestContent, requestType)
	if err != nil || !ok {
		fmt.Printf("File %s not deleted\n", requestContent)
		return
	}

	fmt.Printf("The file %s deleted successfully\n", requestContent)
}

func handleList(client *clientsocket.ClientSocket) {
	list, err := client.ListRequest()
	if err != nil {
		fmt.Println("Files not available")
		return
	}

	fmt.Println("Files list:")
	for id, file := range list {
		fmt.Printf("%d. %s\n", id, file)
	}
}

func cursor(message string) string {
	fmt.Print(message)
	fmt.Print("> ")
	scanner.Scan()
	return scanner.Text()
}
