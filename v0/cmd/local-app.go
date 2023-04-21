package main

import (
	"bufio"
	"fmt"
	"os"

	"fileserver/internal/storage"
)

func main() {
	localStorage := storage.New()

	help()
	localStorage.Help()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">  ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		data, err := localStorage.ExecCommandFromInput(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("success")

		for i, row := range data {
			fmt.Printf("%d. %s\n", i+1, row)
		}
	}
}

func help() {
	fmt.Println("commands:")

	fmt.Print("\texit\t")
	fmt.Println(" - end the programm")
}
