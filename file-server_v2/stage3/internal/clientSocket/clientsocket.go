package clientsocket

import (
	"encoding/json"
	"file_server/internal/models"
	"fmt"
	"net"
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

func (c ClientSocket) GetRequest(filename string) (string, error) {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return "", err
	}

	request := models.RequestBody{
		Command:  "get",
		Filename: filename,
	}

	response, err := send(connection, request)
	if err != nil {
		return "", err
	}

	responseData := models.ResponseGetBody{}
	err = json.Unmarshal(response, &responseData)
	if err != nil {
		return "", err
	}

	if !responseData.Ok {
		return "", fmt.Errorf("file %s unavailable", filename)
	}

	return responseData.Content, nil
}

func (c ClientSocket) CreateRequest(filename, fileContent string) (bool, error) {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return false, err
	}

	request := models.RequestBody{
		Command:  "add",
		Filename: filename,
		Content:  fileContent,
	}

	response, err := send(connection, request)
	if err != nil {
		return false, err
	}

	fmt.Println("The request was sent.")

	responseData := models.ResponseBody{}
	err = json.Unmarshal(response, &responseData)
	if err != nil {
		return false, err
	}

	return responseData.Ok, nil
}

func (c ClientSocket) DeleteRequest(filename string) (bool, error) {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return false, err
	}

	request := models.RequestBody{
		Command:  "delete",
		Filename: filename,
	}

	response, err := send(connection, request)
	if err != nil {
		return false, err
	}

	fmt.Println("The request was sent.")

	responseData := models.ResponseBody{}
	err = json.Unmarshal(response, &responseData)
	if err != nil {
		return false, err
	}

	return responseData.Ok, nil
}

func (c ClientSocket) ListRequest() ([]string, error) {
	connection, err := net.Dial("tcp", c.Address)
	if err != nil {
		return nil, err
	}

	request := models.RequestBody{
		Command: "list",
	}

	response, err := send(connection, request)
	if err != nil {
		return nil, err
	}

	fmt.Println("The request was sent.")

	responseData := models.ResponseListBody{}
	err = json.Unmarshal(response, &responseData)
	if err != nil {
		return nil, err
	}

	if !responseData.Ok {
		return nil, fmt.Errorf("files unavailable")
	}

	return responseData.List, nil
}

func send(connection net.Conn, request models.RequestBody) ([]byte, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	_, err = connection.Write(requestBytes)
	if err != nil {
		return nil, err
	}

	inputStream := make([]byte, 1024)
	length, err := connection.Read(inputStream)
	if err != nil {
		return nil, err
	}

	return inputStream[:length], nil
}
