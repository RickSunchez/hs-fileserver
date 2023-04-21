// console_storage obj function

package main

import (
	"file_server/internal/consolestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	consoleStorage := consolestorage.New()

	consoleStorage.Add("file1.txt")
	consoleStorage.Add("file2.pdf")
	consoleStorage.Add("file3.log")
	consoleStorage.Add("fileToDel.del")

	router := gin.Default()

	router.POST("/add", func(context *gin.Context) {
		context.String(http.StatusOK, "file added")
	})
	router.DELETE("/delete", func(context *gin.Context) {
		context.String(http.StatusOK, "file deleted")
	})
	router.GET("/get", func(context *gin.Context) {
		context.String(http.StatusOK, "file file transfered")
	})
	router.GET("/list", func(context *gin.Context) {
		context.String(http.StatusOK, "list of files")
	})

	router.Run()
}
