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
var helloMessage = "Enter action\n  1 - get a file, 2 - create a file, 3 - delete a file, 4 - list files, 5 - exit:\n"

func main() {
	var input string
	client := clientsocket.New(address)

	err := client.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Client started!")

	for {
		cursor(helloMessage)
		input = scanner.Text()

		switch input {
		case "1": // get
			cursor("Enter filename:\n")

			filename := scanner.Text()

			content, err := client.GetRequest(filename)
			if err != nil {
				fmt.Printf("File %s unavailable\n", filename)
				continue
			}

			fmt.Printf("File %s content:\n", filename)
			fmt.Println(content + "\n")

		case "2": // create
			cursor("Enter filename:\n")
			filename := scanner.Text()

			cursor("Enter file content:\n")
			content := scanner.Text()

			ok, err := client.CreateRequest(filename, content)
			if err != nil || !ok {
				fmt.Printf("File %s not created\n", filename)
				continue
			}

			fmt.Printf("The file %s added successfully\n", filename)

		case "3": // delete
			cursor("Enter filename:\n")

			filename := scanner.Text()

			ok, err := client.DeleteRequest(filename)
			if err != nil || !ok {
				fmt.Printf("File %s not deleted\n", filename)
				continue
			}

			fmt.Printf("The file %s deleted successfully\n", filename)
		case "4": // list
			list, err := client.ListRequest()
			if err != nil {
				fmt.Println("Files not available")
				continue
			}

			fmt.Println("Files list:")
			for i, file := range list {
				fmt.Printf("%d. %s\n", i+1, file)
			}
		case "5": // exit
			os.Exit(0)
		default:
			fmt.Println("Undefined command")
			continue
		}
	}

}

func cursor(message string) {
	fmt.Print(message)
	fmt.Print("> ")
	scanner.Scan()
}
