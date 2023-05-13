package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	DataFolder    = "./data/"
	AddCommand    = "add"
	GetCommand    = "get"
	DeleteCommand = "delete"
	ListCommand   = "list"
)

func main() {
	consoleStorage := NewLocalStorage()

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
		case AddCommand:
			err := consoleStorage.Add(filename)
			if err != nil {
				fmt.Printf("Cannot add the file %s\n", filename)
				continue
			}
			fmt.Printf("The file %s added successfully\n", filename)
		case GetCommand:
			_, err := consoleStorage.Get(filename)
			if err != nil {
				fmt.Printf("The file %s not found\n", filename)
				continue
			}
			fmt.Printf("The file %s was sent\n", filename)
		case DeleteCommand:
			err := consoleStorage.Delete(filename)
			if err != nil {
				fmt.Printf("The file %s not found\n", filename)
				continue
			}
			fmt.Printf("The file %s was deleted\n", filename)
		case ListCommand:
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

type LocalStorage struct{}

func NewLocalStorage() LocalStorage {
	return LocalStorage{}
}

func (ls LocalStorage) Get(filename string) (string, error) {
	path := filepath.Join(DataFolder, filename)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}

	return path, nil
}

func (ls LocalStorage) List() ([]string, error) {
	files, err := ioutil.ReadDir(DataFolder)
	if err != nil {
		return nil, err
	}

	availableFiles := []string{}
	for _, file := range files {
		if !file.IsDir() {
			availableFiles = append(availableFiles, file.Name())
		}
	}

	return availableFiles, nil
}

func (ls LocalStorage) Delete(filename string) error {
	path := filepath.Join(DataFolder, filename)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func (ls LocalStorage) Add(filename string) error {
	path := filepath.Join(DataFolder, filename)
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
