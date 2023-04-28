package clientsocket

import (
	"encoding/json"
	"errors"
	"file_server/internal/models"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
)

const (
	DataFolder = "./data/client/"
)

type ClientSocket struct {
	Address string
}

func New(address string) ClientSocket {
	return ClientSocket{
		Address: address,
	}
}

func (c ClientSocket) Ping() error {
	_, err := net.Dial("tcp", c.Address)
	if err != nil {
		return err
	}

	return nil
}

func (c ClientSocket) GetRequest(requestContent, requestType string) error {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return err
	}

	request := models.RequestBody{
		Command: "get",
		Type:    requestType,
		FileId:  requestContent,
	}

	response := models.ResponseGetBody{}
	err = send(connection, request, &response)
	if err != nil {
		return err
	}

	if !response.Ok {
		return fmt.Errorf("file %s unavailable", requestContent)
	}

	path := filepath.Join(DataFolder, response.Filename)
	err = ioutil.WriteFile(path, response.Content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c ClientSocket) SendRequest(filename string) (uint32, error) {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return 0, err
	}

	path := filepath.Join(DataFolder, filename)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	request := models.RequestBody{
		Command:  "add",
		Filename: filename,
		Content:  fileContent,
	}

	response := models.ResponseAddBody{}
	err = send(connection, request, &response)
	if err != nil {
		return 0, err
	}

	if !response.Ok {
		return 0, errors.New("file not created")
	}

	return response.FileId, nil
}

func (c ClientSocket) DeleteRequest(requestContent, requestType string) (bool, error) {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return false, err
	}

	request := models.RequestBody{
		Command: "delete",
		Type:    requestType,
		FileId:  requestContent,
	}

	response := models.ResponseBody{}
	err = send(connection, request, &response)
	if err != nil {
		return false, err
	}

	return response.Ok, nil
}

func (c ClientSocket) ListRequest() (map[uint32]string, error) {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return nil, err
	}

	request := models.RequestBody{
		Command: "list",
	}

	response := models.ResponseListBody{}
	err = send(connection, request, &response)
	if err != nil {
		return nil, err
	}

	if !response.Ok {
		return nil, fmt.Errorf("files unavailable")
	}

	return response.List, nil
}

func send(connection net.Conn, request, response any) error {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = connection.Write(requestBytes)
	if err != nil {
		return err
	}

	fmt.Println("The request was sent.")

	inputStream := make([]byte, 10*1024*1024*1024) // limit 10 Mb
	length, err := connection.Read(inputStream)
	if err != nil {
		return err
	}

	err = json.Unmarshal(inputStream[:length], response)
	if err != nil {
		return err
	}

	return nil
}
