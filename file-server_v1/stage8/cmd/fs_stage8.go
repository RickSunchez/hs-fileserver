// console_storage obj function

package main

import (
	"file_server/internal/httphelper"

	"github.com/gin-gonic/gin"
)

func main() {
	// consoleStorage := consolestorage.New()

	// consoleStorage.Add("file1.txt")
	// consoleStorage.Add("file2.pdf")
	// consoleStorage.Add("file3.log")
	// consoleStorage.Add("fileToDel.del")

	router := gin.Default()

	router.POST("/add", httphelper.HandleAddRequest)
	router.DELETE("/delete", httphelper.HandleDeleteRequest)
	router.GET("/get", httphelper.HandleGetRequest)
	router.GET("/list", httphelper.HandleListRequest)

	router.Run()
}
