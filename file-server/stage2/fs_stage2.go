// delete file function

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
	Add("file1.txt")
	Add("file2.pdf")
	Add("file3.log")

	var filename string
	fmt.Println("Enter filename to delete:")
	fmt.Scanln(&filename)

	err := Delete(filename)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println("success")
}

func Delete(filename string) error {
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
