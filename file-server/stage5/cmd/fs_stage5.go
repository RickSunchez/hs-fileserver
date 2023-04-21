// console_storage obj function

package main

import (
	"console_storage/internal/consolestorage"
	"fmt"
)

func main() {
	consoleStorage := consolestorage.New()

	consoleStorage.Add("file1.txt")
	consoleStorage.Add("file2.pdf")
	consoleStorage.Add("file3.log")
	consoleStorage.Add("fileToDel.del")

	list, err := consoleStorage.List()
	if err != nil {
		fmt.Println("error to get files list")
		return
	}

	fmt.Println("Init list: ")
	for _, filename := range list {
		fmt.Println(filename)
	}

	fmt.Println("Delete existing file")
	err = consoleStorage.Delete("fileToDel.del")
	if err != nil {
		fmt.Println("error to delete file")
		return
	} else {
		fmt.Println("success")
	}

	fmt.Println("Get existing file")
	filePath, err := consoleStorage.Get("file3.log")
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(filePath)
	fmt.Println("success")

	list, err = consoleStorage.List()
	if err != nil {
		fmt.Println("error to get files list")
		return
	}

	fmt.Println("Final list: ")
	for _, filename := range list {
		fmt.Println(filename)
	}
}
