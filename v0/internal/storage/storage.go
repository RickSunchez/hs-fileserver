package storage

import (
	"errors"
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
	ListCommand   = "list"
	DeleteCommand = "delete"
)

type LocalStorage struct{}

func New() LocalStorage {
	return LocalStorage{}
}

func (ls LocalStorage) Add(filenames ...string) error {
	for _, filename := range filenames {
		file, err := os.Create(filepath.Join(DataFolder, filename))
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}

func (ls LocalStorage) Get(filenames ...string) ([]string, error) {
	filePaths := []string{}
	for _, filename := range filenames {
		filePath := filepath.Join(DataFolder, filename)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			fmt.Printf("file \"%s\" not found\n", filename)
			continue
		}

		fmt.Printf("file \"%s\" was sent\n", filename)
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}

func (ls LocalStorage) List() ([]string, error) {
	files, err := ioutil.ReadDir(DataFolder)
	if err != nil {
		return []string{}, err
	}

	availableFiles := []string{}
	for _, file := range files {
		if !file.IsDir() {
			availableFiles = append(availableFiles, file.Name())
		}
	}

	return availableFiles, nil
}

func (ls LocalStorage) Delete(filenames ...string) error {
	for _, filename := range filenames {
		filePath := filepath.Join(DataFolder, filename)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			fmt.Printf("file \"%s\" not found\n", filename)
			continue
		}

		// @todo add file read and sent
		err = os.Remove(filePath)
		if err != nil {
			fmt.Printf("file \"%s\" can't be removed", filename)
			continue
		}
		fmt.Printf("file \"%s\" was deleted\n", filename)
	}

	return nil
}

func (ls LocalStorage) ExecCommandFromInput(inputData string) ([]string, error) {
	args := strings.Fields(inputData)
	if len(args) == 0 {
		return []string{}, errors.New("invalid input")
	}

	command := args[0]
	params := args[1:]

	switch command {
	case AddCommand:
		if len(params) == 0 {
			return nil, errors.New("add command requires at least 1 argument")
		}
		return nil, ls.Add(params...)
	case GetCommand:
		if len(params) == 0 {
			return nil, errors.New("get command requires at least 1 argument")
		}
		return ls.Get(params...)
	case ListCommand:
		return ls.List()
	case DeleteCommand:
		if len(params) == 0 {
			return nil, errors.New("delete command requires at least 1 argument")
		}
		return nil, ls.Delete(params...)
	default:
		return nil, errors.New("undefined command")
	}
}

func (ls LocalStorage) Help() {
	fmt.Println("storage commands:")

	fmt.Print("\tadd <filename_1> <filename_2> ...\t")
	fmt.Println(" - produces files that have specified names (<filename_1> <filename_2> etc.)")

	fmt.Print("\tget <filename_1> <filename_2> ...\t")
	fmt.Println(" - sent files by specified filenames (if files exist)")

	fmt.Print("\tlist\t\t\t\t\t")
	fmt.Println(" - return a list of available files")

	fmt.Print("\tdelete  <filename_1> <filename_2> ...\t")
	fmt.Println(" - delete files that have specified names (<filename_1> <filename_2> etc.)")
}
