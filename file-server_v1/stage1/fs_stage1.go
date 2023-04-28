// create file function

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	DataFolder = "./data/"
)

func main() {
	var filename string

	fmt.Println("Enter filename:")
	fmt.Scanln(&filename)

	err := Add(filename)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println("success")
}

func Add(filename string) error {
	if !filenameIsSuitable(filename) {
		return errors.New("incorrect filename")
	}

	path := filepath.Join(DataFolder, filename)
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func filenameIsSuitable(filename string) bool {
	rejectedSynbols := []string{":", "»", "\\", "/", "?", "|", " "}

	for _, symbol := range rejectedSynbols {
		if strings.Contains(filename, symbol) {
			return false
		}
	}

	return true
}
