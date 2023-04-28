package serversocket

import (
	"encoding/json"
	"errors"
	"file_server/internal/consolestorage"
	"file_server/internal/models"
	"log"
	"net"
	"os"
)

type ServerSocket struct {
	Address string
}

const (
	GetCommand    = "get"
	AddCommand    = "add"
	DeleteCommand = "delete"
	ListCommand   = "list"
)

var storage = consolestorage.New()

func New(address string) ServerSocket {
	return ServerSocket{
		Address: address,
	}
}

func (s ServerSocket) Listen() {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Server started at %s\n", s.Address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	request, err := inputData(connection)
	if err != nil {
		log.Println(err)
		return
	}

	var response any
	err = nil

	switch request.Command {
	case GetCommand:
		response, err = handleGetCommand(request)
	case AddCommand:
		response, err = handleAddCommand(request)
	case DeleteCommand:
		response, err = handleDeleteCommand(request)
	case ListCommand:
		response, err = handleListCommand(request)
	default:
		err = errors.New("undefined command: " + request.Command)
		response = models.ResponseBody{
			Ok: false,
		}
	}

	if err != nil {
		log.Println(err)
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	connection.Write(responseBytes)
	log.Printf("Sended: \n%v\n", response)
}

func inputData(connection net.Conn) (*models.RequestBody, error) {
	inputStream := make([]byte, 1024)
	length, err := connection.Read(inputStream)
	if err != nil {
		return nil, err
	}

	request := models.RequestBody{}
	err = json.Unmarshal(inputStream[:length], &request)
	if err != nil {
		return nil, err
	}

	log.Printf("Received: \n%v\n", request)

	return &request, nil
}

/* Commands handlers */

func handleGetCommand(request *models.RequestBody) (*models.ResponseGetBody, error) {
	response := models.ResponseGetBody{
		Ok: false,
	}
	filePath, err := storage.Get(request.Filename)
	if err != nil {
		return &response, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return &response, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return &response, err
	}
	fileSize := fileInfo.Size()

	fileContent := make([]byte, fileSize)
	_, err = file.Read(fileContent)
	if err != nil {
		return &response, err
	}

	response.Ok = true
	response.Content = string(fileContent)

	return &response, nil
}

func handleAddCommand(request *models.RequestBody) (*models.ResponseBody, error) {
	response := models.ResponseBody{
		Ok: false,
	}

	err := storage.Add(request.Filename, request.Content)
	if err != nil {
		return &response, err
	}

	response.Ok = true
	return &response, nil
}

func handleDeleteCommand(request *models.RequestBody) (*models.ResponseBody, error) {
	response := models.ResponseBody{
		Ok: false,
	}

	err := storage.Delete(request.Filename)
	if err != nil {
		return &response, err
	}

	response.Ok = true
	return &response, nil
}

func handleListCommand(request *models.RequestBody) (*models.ResponseListBody, error) {
	response := models.ResponseListBody{
		Ok: false,
	}

	list, err := storage.List()
	if err != nil {
		return &response, err
	}

	response.Ok = true
	response.List = list
	return &response, nil
}
