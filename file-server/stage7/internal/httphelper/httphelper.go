package httphelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageAnswer struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func HandleAddRequest(context *gin.Context) {
	response := MessageAnswer{
		Ok:      true,
		Message: "file created",
	}
	context.JSON(http.StatusCreated, response)
}

func HandleDeleteRequest(context *gin.Context) {
	response := MessageAnswer{
		Ok:      true,
		Message: "file deleted",
	}
	context.JSON(http.StatusOK, response)
}

func HandleGetRequest(context *gin.Context) {
	response := MessageAnswer{
		Ok:      true,
		Message: "file file transfered",
	}
	context.JSON(http.StatusOK, response)
}

func HandleListRequest(context *gin.Context) {
	response := MessageAnswer{
		Ok:      true,
		Message: "list of files",
	}
	context.JSON(http.StatusOK, response)
}
