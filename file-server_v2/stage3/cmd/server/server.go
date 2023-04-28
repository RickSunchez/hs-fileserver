package main

import serversocket "file_server/internal/serverSocket"

func main() {
	server := serversocket.New("127.0.0.1:23456")
	server.Listen()
}
