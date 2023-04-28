// console_storage obj function

package main

import (
	"bufio"
	"file_server/internal/consolestorage"
	"fmt"
	"os"
	"strings"
)

func main() {
	consoleStorage := consolestorage.New()

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print(">  ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]

		var filename string
		if len(args) >= 2 {
			filename = args[1]
		}

		switch command {
		case consolestorage.AddCommand:
			err := consoleStorage.Add(filename)
			if err != nil {
				fmt.Printf("Cannot add the file %s\n", filename)
				continue
			}
			fmt.Printf("The file %s added successfully\n", filename)
		case consolestorage.GetCommand:
			_, err := consoleStorage.Get(filename)
			if err != nil {
				fmt.Printf("The file %s not found\n", filename)
				continue
			}
			fmt.Printf("The file %s was sent\n", filename)
		case consolestorage.DeleteCommand:
			err := consoleStorage.Delete(filename)
			if err != nil {
				fmt.Printf("The file %s not found\n", filename)
				continue
			}
			fmt.Printf("The file %s was deleted\n", filename)
		case consolestorage.ListCommand:
			list, err := consoleStorage.List()
			if err != nil {
				fmt.Printf("Storage is empty\n")
				continue
			}
			fmt.Printf("Storaged files:\n")
			for i, file := range list {
				fmt.Printf("%d. %s\n", i+1, file)
			}
		}
	}
}
