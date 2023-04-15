package main

import (
	"fileserver/internal/httphelper"
	"fileserver/internal/storage"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	localStorage := storage.New()

	router.GET("/add", func(context *gin.Context) {
		filesQuery := context.Query("files")
		filesNames := strings.Split(filesQuery, ",")

		err := localStorage.Add(filesNames...)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				httphelper.Response{"invalid request"},
			)
			return
		}

		context.JSON(
			http.StatusOK,
			httphelper.Response{"success"},
		)
	})

	router.GET("/get", func(context *gin.Context) {
		file := context.Query("file")

		filePath, err := localStorage.Get(file)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				httphelper.Response{"invalid request"},
			)
			return
		}

		context.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath[0]))
		context.File(filePath[0])
	})

	router.GET("/list", func(context *gin.Context) {
		list, err := localStorage.List()
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				httphelper.Response{"invalid request"},
			)
			return
		}

		context.JSON(
			http.StatusOK,
			httphelper.ListResponse{list},
		)
	})
	router.GET("/delete", func(context *gin.Context) {
		filesQuery := context.Query("files")
		filesNames := strings.Split(filesQuery, ",")

		err := localStorage.Delete(filesNames...)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				httphelper.Response{"invalid request"},
			)
			return
		}

		context.JSON(
			http.StatusOK,
			httphelper.Response{"success"},
		)
	})

	router.Run()
}
