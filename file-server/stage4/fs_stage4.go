// get file function

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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

	fmt.Println("Enter filename to get:")
	fmt.Scanln(&filename)

	filePath, err := Get(filename)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(filePath)
	fmt.Println("success")
}

func Get(filename string) (string, error) {
	path := filepath.Join(DataFolder, filename)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}

	return path, nil
}

func List() ([]string, error) {
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
