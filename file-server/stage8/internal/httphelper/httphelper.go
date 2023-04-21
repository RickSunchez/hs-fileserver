package httphelper

import (
	"fmt"
	"net/http"
	"path/filepath"

	"file_server/internal/consolestorage"

	"github.com/gin-gonic/gin"
)

type MessageResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type FilesListResponse struct {
	Ok   bool     `json:"ok"`
	List []string `json:"list"`
}

var consoleStorage = consolestorage.New()

func HandleAddRequest(context *gin.Context) {
	filename := context.Query("filename")

	err := consoleStorage.Add(filename)
	if err != nil {
		response := MessageResponse{
			Ok:      false,
			Message: fmt.Sprintf("can't create file %s", filename),
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := MessageResponse{
		Ok:      true,
		Message: fmt.Sprintf("file %s created", filename),
	}
	context.JSON(http.StatusCreated, response)
}

func HandleDeleteRequest(context *gin.Context) {
	filename := context.Query("filename")

	err := consoleStorage.Delete(filename)
	if err != nil {
		response := MessageResponse{
			Ok:      false,
			Message: fmt.Sprintf("can't delete file %s", filename),
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := MessageResponse{
		Ok:      true,
		Message: fmt.Sprintf("file %s deleted", filename),
	}
	context.JSON(http.StatusOK, response)
}

func HandleGetRequest(context *gin.Context) {
	filename := context.Query("filename")

	filePath, err := consoleStorage.Get(filename)
	if err != nil {
		response := MessageResponse{
			Ok:      false,
			Message: fmt.Sprintf("can't find file %s", filename),
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	context.Status(http.StatusOK)
	context.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	context.File(filePath)
}

func HandleListRequest(context *gin.Context) {
	list, err := consoleStorage.List()
	if err != nil {
		response := MessageResponse{
			Ok:      false,
			Message: "files not found",
		}
		context.JSON(http.StatusNoContent, response)
		return
	}

	response := FilesListResponse{
		Ok:   true,
		List: list,
	}
	context.JSON(http.StatusOK, response)
}
