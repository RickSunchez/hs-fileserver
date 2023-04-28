package consolestorage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	DataFolder = "./data/server/"
	IndexFile  = "./data/server/index.json"
)

type IndexStruct struct {
	Files          map[uint32]string `json:"files"`
	LastInsertedId uint32            `json:"last_inserted_id"`
}

type LocalStorage struct{}

func New() LocalStorage {
	_, err := os.Stat(IndexFile)
	if os.IsNotExist(err) {
		file, err := os.Create(IndexFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		emptyIndex := IndexStruct{
			Files:          map[uint32]string{},
			LastInsertedId: 0,
		}
		err = writeIndex(&emptyIndex)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return LocalStorage{}
}

func (ls LocalStorage) Get(fileId, requestType string) (string, error) {
	var filename string
	switch requestType {
	case "1": // name
		filename = fileId
	case "2": // id
		index, err := readIndex()
		if err != nil {
			return "", err
		}

		fileIdInt, err := strconv.Atoi(fileId)
		if err != nil {
			return "", err
		}

		filename = index.Files[uint32(fileIdInt)]
	default:
		return "", errors.New("undefined request type")
	}

	path := filepath.Join(DataFolder, filename)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}

	return path, nil
}

func (ls LocalStorage) List() (map[uint32]string, error) {
	index, err := readIndex()
	if err != nil {
		return nil, err
	}

	return index.Files, nil
}

func (ls LocalStorage) Delete(fileId, requestType string) error {
	var filename string
	var fileId32 uint32
	index, err := readIndex()
	if err != nil {
		return err
	}

	switch requestType {
	case "1": // name
		filename = fileId
		for fileId, name := range index.Files {
			if name == filename {
				fileId32 = fileId
				break
			}
		}

	case "2": // id
		fileIdInt, err := strconv.Atoi(fileId)
		if err != nil {
			return err
		}

		fileId32 = uint32(fileIdInt)
		filename = index.Files[fileId32]
	default:
		return errors.New("undefined request type")
	}

	path := filepath.Join(DataFolder, filename)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	delete(index.Files, fileId32)
	err = writeIndex(index)
	if err != nil {
		return err
	}

	return nil
}

func (ls LocalStorage) Add(filename string, fileContent []byte) (uint32, error) {
	if !filenameIsSuitable(filename) {
		return 0, errors.New("incorrect filename")
	}

	path := filepath.Join(DataFolder, filename)

	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	_, err = file.Write(fileContent)
	if err != nil {
		return 0, err
	}

	fileId, err := indexNewFile(filename)
	if err != nil {
		return 0, err
	}

	return fileId, nil
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

func indexNewFile(filename string) (uint32, error) {
	index, err := readIndex()
	if err != nil {
		return 0, err
	}

	index.LastInsertedId += 1
	index.Files[index.LastInsertedId] = filename

	return index.LastInsertedId, writeIndex(index)
}

func writeIndex(data *IndexStruct) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(IndexFile, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func readIndex() (*IndexStruct, error) {
	indexBytes, err := ioutil.ReadFile(IndexFile)
	if err != nil {
		return nil, err
	}

	var index IndexStruct
	err = json.Unmarshal(indexBytes, &index)
	if err != nil {
		return nil, err
	}

	return &index, nil
}
